package httpsrv

import (
	"context"
	"errors"
	"net/http"

	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/api"
	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/db"
	"github.com/romanchechyotkin/effective-mobile-test-task/schema"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "http_server"

func NewModule() fx.Option {
	return fx.Module(
		ModuleName,

		fx.Provide(NewServer, NewConfig),

		fx.Options(api.NewModule(), db.NewModule()),

		fx.Invoke(func(
			lc fx.Lifecycle,
			log *zap.Logger,
			server *Server,
			dbCfg *db.Config,
		) {
			lc.Append(
				fx.Hook{
					OnStart: func(_ context.Context) error {
						db.Migrate(log, &schema.DB, dbCfg.URL)

						go func() {
							log.Info("http-server listen and serve", zap.String("address", server.base.Addr))
							if err := server.base.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
								log.Error("failed to listen and serve", zap.Error(err))
							}
						}()

						return nil
					},
					OnStop: func(ctx context.Context) error {
						if err := server.base.Shutdown(ctx); err != nil {
							log.Error("failed to shutdown http server", zap.Error(err))
							return err
						}

						return nil
					},
				})
		}),

		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named(ModuleName)
		}),
	)
}
