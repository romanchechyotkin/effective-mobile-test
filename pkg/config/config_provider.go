package config

import (
	"embed"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/config"
)

const (
	AppConfigFile = "config.yaml"
)

func NewConfigProvider(configFs *embed.FS) func() (*config.YAML, error) {
	return func() (*config.YAML, error) {
		fileYmlOpt, err := GetFileYmlOpt(configFs)
		if err != nil {
			return nil, err
		}

		sources := make([]config.YAMLOption, 0, 2)
		sources = append(sources, fileYmlOpt)

		return config.NewYAML(sources...)
	}
}

func GetFileYmlOpt(configFs *embed.FS) (config.YAMLOption, error) {
	file, err := configFs.Open(AppConfigFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return config.RawSource(file), nil
}
