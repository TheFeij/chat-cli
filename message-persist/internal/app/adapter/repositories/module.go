package repositories

import (
	"go.uber.org/fx"
	"message-persist/internal/app/adapter/repositories/mongodb"
)

var Module = fx.Options(
	mongodb.Module,
)
