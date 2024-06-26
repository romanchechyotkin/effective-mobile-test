package httpsrv

import (
	"fmt"

	"go.uber.org/config"
)

type HTTPConfig struct {
	Port string `yaml:"port"`
	Bind string `yaml:"bind"`
}

func NewConfig(provider *config.YAML) (*HTTPConfig, error) {
	var cfg HTTPConfig
	var err error

	if err := provider.Get("http").Populate(&cfg); err != nil {
		err = fmt.Errorf("failed to get http config: %w", err)
	}

	return &cfg, err
}
