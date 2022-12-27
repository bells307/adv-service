package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/bells307/adv-service/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bells307/adv-service/internal/domain"
)

const CAT_COLLECTION_NAME = "category"

type categoryMongoDBRepository struct {
	client *mongodb.MongoDBClient
}

func NewCategoryMongoDBRepository(client *mongodb.MongoDBClient) *categoryMongoDBRepository {
	return &categoryMongoDBRepository{client}
}

func (r *categoryMongoDBRepository) Collection() *mongo.Collection {
	return r.client.Collection(CAT_COLLECTION_NAME)
}

func (r *categoryMongoDBRepository) Create(ctx context.Context, category domain.Category) error {
	_, err := r.Collection().InsertOne(ctx, category)
	return err
}

func (r *categoryMongoDBRepository) Delete(ctx context.Context, id string) error {
	_, err := r.Collection().DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *categoryMongoDBRepository) FindByID(ctx context.Context, id string) (domain.Category, error) {
	res := r.Collection().FindOne(ctx, bson.M{"_id": id})
	err := res.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Category{}, domain.ErrCategoryNotFound
		} else {
			return domain.Category{}, err
		}
	}

	var category domain.Category
	res.Decode(&category)

	return category, nil
}

func (r *categoryMongoDBRepository) FindAllByID(ctx context.Context, ids []string) ([]domain.Category, error) {
	bsonArr := bson.A{}
	for _, id := range ids {
		bsonArr = append(bsonArr, id)
	}

	cur, err := r.Collection().Find(ctx, bson.M{
		"_id": bson.M{"$in": bsonArr},
	})

	if err != nil {
		return []domain.Category{}, fmt.Errorf("error finding category in mongodb: %v", err)
	}
	defer cur.Close(ctx)

	var categories []domain.Category
	if err := cur.All(ctx, &categories); err != nil {
		return []domain.Category{}, fmt.Errorf(
			"error decoding mongodb cursor while finding category in mongodb: %v",
			err,
		)
	}

	return categories, nil
}
