package api

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "api client"

func NewModule() fx.Option {
	return fx.Module(
		ModuleName,

		fx.Provide(NewClient, NewConfig),

		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named(ModuleName)
		}),
	)
}
