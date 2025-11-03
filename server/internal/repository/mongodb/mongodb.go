package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"time"
)

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient() *MongoClient {
	client, err := mongo.Connect(
		options.Client().
			ApplyURI("mongodb://localhost:27017").
			SetAuth(options.Credential{
				AuthSource: "admin",
				Username:   "root",
				Password:   "1234",
			}))
	if err != nil {
		panic(err)
	}

	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")

	return &MongoClient{
		client: client,
	}
}
