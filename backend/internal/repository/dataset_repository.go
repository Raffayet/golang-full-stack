package repository

import (
	"context"
	"data-merging/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatasetRepository interface {
	FetchDatasets(ctx context.Context) ([]domain.Dataset, error)
	Client() *mongo.Client // Expose MongoDB client
}

type MongoDatasetRepository struct {
	client *mongo.Client
}

func NewDatasetRepository(client *mongo.Client) *MongoDatasetRepository {
	return &MongoDatasetRepository{client: client}
}

// Client returns the MongoDB client from the repository
func (repo *MongoDatasetRepository) Client() *mongo.Client {
	return repo.client
}

func (repo *MongoDatasetRepository) FetchDatasets(ctx context.Context) ([]domain.Dataset, error) {
	var datasets []domain.Dataset
	collection := repo.client.Database("data_merging").Collection("datasets")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var dataset domain.Dataset
		if err := cursor.Decode(&dataset); err != nil {
			return nil, err
		}
		datasets = append(datasets, dataset)
	}

	// Check if there were any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return datasets, nil
}
