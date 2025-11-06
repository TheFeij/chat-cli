package adapter

import (
	"go.uber.org/fx"
	"message-persist/internal/app/adapter/repositories"
)

var Module = fx.Options(
	repositories.Module,
)
