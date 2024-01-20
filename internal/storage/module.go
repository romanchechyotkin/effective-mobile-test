package storage

import (
	"context"

	"github.com/romanchechyotkin/effective-mobile-test-task/internal/storage/repo"
	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/db"
	"github.com/romanchechyotkin/effective-mobile-test-task/schema"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "storage"

func NewModule() fx.Option {
	return fx.Module(
		ModuleName,

		fx.Provide(repo.NewUsers),

		fx.Options(db.NewModule()),

		fx.Invoke(func(
			lc fx.Lifecycle,
			log *zap.Logger,
			dbCfg *db.Config,
		) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					db.Migrate(log, &schema.DB, dbCfg.URL)

					return nil
				},
			})
		}),

		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named(ModuleName)
		}),
	)
}
