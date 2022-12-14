package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/bells307/adv-service/internal/domain/advertisment/model"
	"github.com/bells307/adv-service/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_NAME = "advertisment"

type advertismentMongoDBRepository struct {
	client *mongodb.MongoDBClient
}

func NewAdvertismentMongoDBRepository(client *mongodb.MongoDBClient) *advertismentMongoDBRepository {
	return &advertismentMongoDBRepository{client}
}

func (r *advertismentMongoDBRepository) Collection() *mongo.Collection {
	return r.client.Collection(COLLECTION_NAME)
}

// Получить объявления
func (r *advertismentMongoDBRepository) Get(ctx context.Context, limit int64, offset int64) ([]*model.Advertisment, error) {
	filter := bson.D{}
	opts := options.Find().SetLimit(limit).SetSkip(offset)

	cur, err := r.Collection().Find(ctx, filter, opts)
	if err != nil {
		return []*model.Advertisment{}, fmt.Errorf("can't find advertisments in mongo collection: %v", err)
	}
	defer cur.Close(ctx)

	var advertisments []*model.Advertisment
	if err := cur.All(ctx, &advertisments); err != nil {
		return []*model.Advertisment{}, fmt.Errorf("error decoding mongodb cursor: %v", err)
	}

	return advertisments, nil
}

// Получить объявление
func (r *advertismentMongoDBRepository) GetOne(ctx context.Context, id string) (*model.Advertisment, error) {
	res := r.Collection().FindOne(ctx, bson.M{"_id": id})
	err := res.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	var advertisment model.Advertisment
	res.Decode(&advertisment)

	return &advertisment, nil
}

// Создать объявление
func (r *advertismentMongoDBRepository) Create(ctx context.Context, adv *model.Advertisment) error {
	_, err := r.Collection().InsertOne(ctx, adv)
	return err
}
