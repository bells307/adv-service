package repository

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/bells307/adv-service/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const ADV_COLLECTION_NAME = "advertisment"

type advertismentMongoDBRepository struct {
	client *mongodb.MongoDBClient
}

// Представление документа объявления в mongodb
type Advertisment struct {
	// Идентификатор
	ID string `bson:"_id"`
	// Имя
	Name string `bson:"name"`
	// Категория
	// В базе мы будем хранить идентификаторы категорий. В дальнейшем, при запросе
	// доменной модели, мы будем аггрегировать их с коллекцией категорий
	Categories []string `bson:"categories"`
	// Описание
	Description string `bson:"description"`
	// Цена
	Price domain.Price `bson:"price"`
	// Ссылка на главное изображение
	MainPhotoURL string `bson:"mainPhotoURL"`
	// Ссылки на дополнительные изображения
	AdditionalPhotoURLs []string `bson:"additionalPhotoURLs"`
}

func NewAdvertismentMongoDBRepository(client *mongodb.MongoDBClient) *advertismentMongoDBRepository {
	return &advertismentMongoDBRepository{client}
}

func (r *advertismentMongoDBRepository) Collection() *mongo.Collection {
	return r.client.Collection(ADV_COLLECTION_NAME)
}

// Получить объявления
func (r *advertismentMongoDBRepository) Find(ctx context.Context, limit uint, offset uint) ([]domain.Advertisment, error) {
	// Аггрегируем с коллекцией категорий по полю "id"
	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         CAT_COLLECTION_NAME,
				"localField":   "categories",
				"foreignField": "_id",
				"as":           "categories",
			},
		},
	}

	if limit > 0 {
		pipeline = append(pipeline, bson.M{"$limit": offset + limit})
	}

	if offset > 0 {
		pipeline = append(pipeline, bson.M{"$skip": offset})
	}

	cur, err := r.Collection().Aggregate(ctx, pipeline)
	if err != nil {
		return []domain.Advertisment{}, fmt.Errorf("error finding advertisments in mongodb: %v", err)
	}
	defer cur.Close(ctx)

	var advs []domain.Advertisment
	if err := cur.All(ctx, &advs); err != nil {
		return []domain.Advertisment{}, fmt.Errorf(
			"error decoding mongodb cursor while finding advertisments in mongodb: %v",
			err,
		)
	}

	if len(advs) == 0 {
		return []domain.Advertisment{}, domain.ErrAdvNotFound
	}

	return advs, nil
}

// Проверить, есть ли объявления с категорией
func (r *advertismentMongoDBRepository) ExistsWithCategory(ctx context.Context, categoryID string) (bool, error) {
	c, err := r.Collection().CountDocuments(ctx, bson.M{
		"category": bson.M{
			"_id": categoryID,
		},
	})
	if err != nil {
		return false, fmt.Errorf("error checking category exists for advertisment: %v", err)
	}

	return c > 0, nil
}

// Получить объявление
func (r *advertismentMongoDBRepository) FindByID(ctx context.Context, id string) (domain.Advertisment, error) {
	// Аггрегируем с коллекцией категорий по полю "id"
	pipeline := bson.A{
		bson.M{"$match": bson.M{"_id": id}},
		bson.M{
			"$lookup": bson.M{
				"from":         CAT_COLLECTION_NAME,
				"localField":   "categories",
				"foreignField": "_id",
				"as":           "categories",
			},
		},
	}
	cur, err := r.Collection().Aggregate(ctx, pipeline)
	if err != nil {
		return domain.Advertisment{}, fmt.Errorf("error finding advertisment by ID in mongodb: %v", err)
	}
	defer cur.Close(ctx)

	var advs []domain.Advertisment
	if err := cur.All(ctx, &advs); err != nil {
		return domain.Advertisment{}, fmt.Errorf(
			"error decoding mongodb cursor while finding advertisment by id in mongodb: %v",
			err,
		)
	}

	if len(advs) == 0 {
		return domain.Advertisment{}, domain.ErrAdvNotFound
	}

	return advs[0], nil
}

// Создать объявление
func (r *advertismentMongoDBRepository) Create(ctx context.Context, adv domain.Advertisment) error {
	var categories []string
	for _, cat := range adv.Categories {
		categories = append(categories, cat.ID)
	}

	docAdv := Advertisment{
		ID:                  adv.ID,
		Name:                adv.Name,
		Categories:          categories,
		Description:         adv.Description,
		Price:               adv.Price,
		MainPhotoURL:        adv.MainPhotoURL,
		AdditionalPhotoURLs: adv.AdditionalPhotoURLs,
	}
	_, err := r.Collection().InsertOne(ctx, docAdv)
	return err
}

// Удалить объявление
func (r *advertismentMongoDBRepository) Delete(ctx context.Context, id string) error {
	_, err := r.Collection().DeleteOne(ctx, bson.M{"_id": id})
	return err
}
