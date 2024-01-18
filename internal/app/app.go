package app

import (
	"context"
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romanchechyotkin/effective-mobile-test-task/internal/config"
	"github.com/romanchechyotkin/effective-mobile-test-task/internal/httpsrv"
	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/logger"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

//go:embed config.yaml
var configFs embed.FS

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			logger.NewLog,
			config.NewConfigProvider(&configFs),
			NewAppConfig,
		),

		fx.Options(
			httpsrv.NewModule(),
		),

		fx.Invoke(func(
			lc fx.Lifecycle,
			server *httpsrv.Server,
		) {
			lc.Append(HttpServerOnStart(server))
		}),

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: log,
			}
		}),
	)
}

func HttpServerOnStart(server *httpsrv.Server) fx.Hook {
	return fx.Hook{
		OnStart: func(ctx context.Context) error {

			server.Router.GET("/status", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "ok\n")
			})

			usersGroup := server.Router.Group("/users")
			usersGroup.GET("/")
			usersGroup.POST("/")

			return nil
		},
	}
}
