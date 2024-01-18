package httpsrv

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const moduleName = "http_server"

func NewModule() fx.Option {
	return fx.Module(
		moduleName,

		fx.Provide(NewServer, NewConfig),

		fx.Invoke(func(server *Server, log *zap.Logger, lc fx.Lifecycle) {
			lc.Append(
				fx.Hook{
					OnStart: func(_ context.Context) error {
						go func() {
							log.Info("http-server listen and serve", zap.String("on", server.base.Addr))
							if err := server.base.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
								log.Error("failed to listen and serve", zap.Error(err))
							}
						}()

						return nil
					},
					OnStop: func(ctx context.Context) error {
						if err := server.base.Shutdown(ctx); err != nil {
							return err
						}

						return nil
					},
				})
		}),

		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named(moduleName)
		}),
	)
}
