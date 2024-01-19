package api

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
	AgeURL         string `yaml:"age_url"`
	GenderURL      string `yaml:"gender_url"`
	NationalityURL string `yaml:"nationality_url"`
}

func NewConfig(provider *config.YAML) (*Config, error) {
	var cfg Config
	var err error

	if err := provider.Get("api").Populate(&cfg); err != nil {
		err = fmt.Errorf("failed to get http config: %w", err)
	}

	return &cfg, err
}
