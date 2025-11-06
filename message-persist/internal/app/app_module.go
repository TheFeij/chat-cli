package app

import (
	"go.uber.org/fx"
	"message-persist/internal/app/adapter"
	"message-persist/internal/app/core"
	"message-persist/internal/app/infrastructure/config"
	"message-persist/internal/app/infrastructure/server"
)

func NewApp() *fx.App {
	return fx.New(
		config.Module,
		server.Module,
		core.Module,
		adapter.Module,
		fx.Invoke(server.StartServer),
	)
}
