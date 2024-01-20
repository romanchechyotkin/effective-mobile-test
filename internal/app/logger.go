package app

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

func NewLogger(options ...zap.Option) (*zap.Logger, error) {
	e := strings.ToUpper(os.Getenv("APP_ENV"))

	switch e {
	case "DEV":
		return zap.NewDevelopment(options...)
	case "PROD":
		return zap.NewProduction(options...)
	}

	return nil, fmt.Errorf("wrong app env")
}
