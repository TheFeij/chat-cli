package core

import (
	"go.uber.org/fx"
	"message-persist/internal/app/core/ports"
	"message-persist/internal/app/core/services"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(services.NewMessagePersistService, fx.As(new(ports.MessagePersistService))),
	),
)
