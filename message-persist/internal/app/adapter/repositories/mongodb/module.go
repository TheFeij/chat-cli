package mongodb

import (
	"go.uber.org/fx"
	"message-persist/internal/app/core/ports"
	"message-persist/internal/app/infrastructure/config"
)

var Module = fx.Options(
	fx.Provide(
		func(config *config.AppConfig) Config {
			conn := config.Mongo.ConnectionString

			return Config{ConnectionString: conn}
		},
		ConnectToMongo,
	),
	fx.Provide(
		fx.Annotate(NewMessageRepository, fx.As(new(ports.MessageRepository))),
	),
)
