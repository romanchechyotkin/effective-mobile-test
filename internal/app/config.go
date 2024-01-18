package app

import (
	"sync"

	"go.uber.org/config"
	"go.uber.org/zap"
)

type Config struct {
	Name string
	Env  string
}

func NewAppConfig(provider *config.YAML, log *zap.Logger) (*Config, error) {
	var (
		once sync.Once
		cfg  Config
		err  error
	)

	once.Do(func() {
		if err := provider.Get("app").Populate(&cfg); err != nil {
			log.Error("failed to get app config", zap.Error(err))
		}
	})

	log.Info("app config", zap.Any("cfg", cfg))

	return &cfg, err
}
