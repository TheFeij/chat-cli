package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/fx"
	"message-persist/internal/app/core/ports"
	"message-persist/internal/app/infrastructure/config"
)

var Module = fx.Options(
	fx.Provide(
		func(config *config.AppConfig) MongoDBConfig {
			conn := config.Mongo.ConnectionString
			if conn == "" && config.Environment == "Development" { // for local development
				conn = "mongodb://localhost:27017"
			}

			return MongoDBConfig{URI: conn}
		},
		func(cfg MongoDBConfig) (*mongo.Client, error) {
			return ConnectToMongo(cfg)
		},
	),
	fx.Provide(
		fx.Annotate(NewMessageRepository, fx.As(new(ports.MessageRepository))),
	),
)
