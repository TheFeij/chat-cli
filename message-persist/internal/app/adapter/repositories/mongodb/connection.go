package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	ConnectionString string
}

func ConnectToMongo(cnfg Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cnfg.ConnectionString)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client, nil
}
