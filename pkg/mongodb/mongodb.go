package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	db *mongo.Database
}

func NewMongoDB(cfg MongoDBConfig) (*MongoDBClient, error) {
	opts := options.Client().ApplyURI(cfg.Uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, fmt.Errorf("can't connect to mongodb: %v", err)
	}

	db := client.Database(cfg.DbName)
	if db == nil {
		return nil, fmt.Errorf("can't create a handle to mongodb database %s", cfg.DbName)
	}

	return &MongoDBClient{
		db,
	}, nil
}

func (m *MongoDBClient) Database() *mongo.Database {
	return m.db
}

func (m *MongoDBClient) Collection(col string) *mongo.Collection {
	return m.db.Collection(col)
}
