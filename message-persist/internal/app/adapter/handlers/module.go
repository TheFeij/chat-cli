package handlers

import (
	"go.uber.org/fx"
	"message-persist/internal/app/adapter/handlers/kafka"
)

var Module = fx.Options(
	kafka.Module,
)
