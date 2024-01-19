package app

import (
	"context"
	"embed"
	"github.com/romanchechyotkin/effective-mobile-test-task/internal/httpsrv"
	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/config"
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

func HttpServerOnStart(server httpsrv.HTTPServer) fx.Hook {
	return fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.RegisterRoutes()
			return nil
		},
	}
}
