package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDBConfig struct {
	URI string
}

func ConnectToMongo(cfg MongoDBConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.URI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client, nil
}
