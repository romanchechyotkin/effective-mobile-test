package config

import (
	"embed"
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/config"
)

const (
	envPrefix     = "APP"
	appConfigFile = "config.yaml"
)

func NewConfigProvider(configFs *embed.FS) func() (*config.YAML, error) {
	return func() (*config.YAML, error) {
		fileYmlOpt, err := GetFileYmlOpt(configFs)
		if err != nil {
			return nil, err
		}

		sources := make([]config.YAMLOption, 0, 3)
		sources = append(sources, fileYmlOpt)
		sources = append(sources, GetEnvYmlOpt(envPrefix)...)

		return config.NewYAML(sources...)
	}
}

// GetFileYmlOpt Waiting for the file name "config.yaml"
func GetFileYmlOpt(configFs *embed.FS) (config.YAMLOption, error) {
	file, err := configFs.Open(appConfigFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return config.RawSource(file), nil
}

func GetEnvYmlOpt(prefix string) []config.YAMLOption {
	sources := make([]config.YAMLOption, 0, 2)
	modules := make(map[string]map[string]any)

	prefixLen := len(prefix) + 1

	for _, eRaw := range os.Environ() {
		if strings.HasPrefix(eRaw, prefix) && len(eRaw) > prefixLen && !strings.HasPrefix(eRaw, prefix+"=") {
			sepIdx := strings.IndexAny(eRaw, "=")
			envName, val := strings.ToLower(eRaw[prefixLen:sepIdx]), eRaw[sepIdx+1:]

			if val != "" {
				modPair := strings.Split(envName, "_")
				lmp := len(modPair)

				if lmp == 1 {
					if modules["app"] == nil {
						modules["app"] = make(map[string]any)
					}

					if isArray(val) {
						parts := strings.Split(val[1:len(val)-1], ",")
						modules["app"][envName] = parts
						continue
					}

					if boolean, err := strconv.ParseBool(strings.ToLower(val)); err == nil {
						modules["app"][envName] = boolean
						continue
					}

					if converted, ok := isInt(val); ok {
						modules["app"][envName] = converted
						continue
					}

					modules["app"][envName] = val
				} else {
					if modules[modPair[0]] == nil {
						modules[modPair[0]] = make(map[string]any)
					}

					modName := envName[len(modPair[0]):]

					if isArray(val) {
						parts := strings.Split(val[1:len(val)-1], ",")

						if strings.HasPrefix(modName, "_") {
							modules[modPair[0]][modName[1:]] = parts
						} else {
							modules[modPair[0]][modName] = parts
						}
						continue
					}

					if strings.HasPrefix(modName, "_") {
						if boolean, err := strconv.ParseBool(strings.ToLower(val)); err == nil {
							modules[modPair[0]][modName[1:]] = boolean
							continue
						}

						if converted, ok := isInt(val); ok {
							modules[modPair[0]][modName[1:]] = converted
							continue
						}

						modules[modPair[0]][modName[1:]] = val
					} else {
						if boolean, err := strconv.ParseBool(strings.ToLower(val)); err == nil {
							modules[modPair[0]][modName[1:]] = boolean
						}

						if converted, ok := isInt(val); ok {
							modules[modPair[0]][modName[1:]] = converted
							continue
						}

						modules[modPair[0]][modName[1:]] = val
					}

				}
			}
		}
	}

	sources = append(sources, config.Static(modules))

	return sources
}

func isArray(val string) bool {
	return strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]")
}

func isInt(val string) (int, bool) {
	converted, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}

	return converted, true
}
