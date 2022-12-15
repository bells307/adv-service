package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/bells307/adv-service/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ADV_COLLECTION_NAME = "advertisment"

type advertismentMongoDBRepository struct {
	client *mongodb.MongoDBClient
}

func NewAdvertismentMongoDBRepository(client *mongodb.MongoDBClient) *advertismentMongoDBRepository {
	return &advertismentMongoDBRepository{client}
}

func (r *advertismentMongoDBRepository) Collection() *mongo.Collection {
	return r.client.Collection(ADV_COLLECTION_NAME)
}

// Получить объявления
func (r *advertismentMongoDBRepository) Find(ctx context.Context, limit uint, offset uint) ([]domain.Advertisment, error) {
	filter := bson.D{}
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))

	cur, err := r.Collection().Find(ctx, filter, opts)
	if err != nil {
		return []domain.Advertisment{}, fmt.Errorf("can't find advertisments in mongo collection: %v", err)
	}
	defer cur.Close(ctx)

	var advertisments []domain.Advertisment
	if err := cur.All(ctx, &advertisments); err != nil {
		return []domain.Advertisment{}, fmt.Errorf("error decoding mongodb cursor: %v", err)
	}

	return advertisments, nil
}

// Получить объявление
func (r *advertismentMongoDBRepository) FindByID(ctx context.Context, id string) (domain.Advertisment, error) {
	res := r.Collection().FindOne(ctx, bson.M{"_id": id})
	err := res.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Advertisment{}, domain.ErrAdvNotFound
		} else {
			return domain.Advertisment{}, err
		}
	}

	var advertisment domain.Advertisment
	res.Decode(&advertisment)

	return advertisment, nil
}

// Создать объявление
func (r *advertismentMongoDBRepository) Create(ctx context.Context, adv domain.Advertisment) error {
	_, err := r.Collection().InsertOne(ctx, adv)
	return err
}
