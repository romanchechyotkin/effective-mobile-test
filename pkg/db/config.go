package db

import (
	"fmt"

	"go.uber.org/config"
)

type Config struct {
	//URL Example postgres://postgres:5432@localhost:5432/nh_templates?sslmode=disable
	URL            string `yaml:"url"`
	ConnectTimeOut int    `yaml:"connect_time_out"`
	AutoCreate     bool   `yaml:"auto_create"`
}

const (
	DefaultConnectTimeOut = 10
	DefaultAutoCreate     = true
)

func NewConfig(provider *config.YAML) (*Config, error) {
	var err error
	cfg := Config{
		ConnectTimeOut: DefaultConnectTimeOut,
		AutoCreate:     DefaultAutoCreate,
	}

	if err := provider.Get("db").Populate(&cfg); err != nil {
		err = fmt.Errorf("failed to get db config: %w", err)
	}

	return &cfg, err
}
