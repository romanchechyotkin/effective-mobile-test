package db

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "db"

func NewModule() fx.Option {
	return fx.Module(
		ModuleName,

		fx.Provide(NewConfig, NewQBuilder),

		fx.Invoke(func(q *QBuilder, lc fx.Lifecycle) {
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						return q.Pool().Ping(ctx)
					},
					OnStop: func(_ context.Context) error {
						q.Pool().Close()
						return nil
					},
				},
			)
		}),

		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named(ModuleName)
		}),
	)
}
