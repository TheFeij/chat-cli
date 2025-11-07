package server

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"message-persist/internal/app/infrastructure/config"
	"net"
	"net/http"
	"time"
)

var Module = fx.Options(
	fx.Provide(NewGinServer),
)

func NewGinServer() *gin.Engine {
	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowAllOrigins: true,
	}))

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	return engine
}

func StartServer(lifecycle fx.Lifecycle, cnfg *config.AppConfig, engine *gin.Engine, lc fx.Shutdowner) {
	srv := &http.Server{
		Addr:    net.JoinHostPort(cnfg.ServerAddress, cnfg.ServerPort),
		Handler: engine,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					// Trigger Fx shutdown gracefully
					_ = lc.Shutdown(fx.ExitCode(1))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return srv.Shutdown(ctx)
		},
	})
}
