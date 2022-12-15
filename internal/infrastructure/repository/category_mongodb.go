package repository

import (
	"context"
	"errors"

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

func (r *categoryMongoDBRepository) CreateCategory(ctx context.Context, category domain.Category) error {
	_, err := r.Collection().InsertOne(ctx, category)
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
