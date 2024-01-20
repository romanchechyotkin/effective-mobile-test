package api

import (
	"net/http"

	"go.uber.org/zap"
)

type Client struct {
	cfg    *Config
	log    *zap.Logger
	client *http.Client
}

func NewClient(cfg *Config, logger *zap.Logger) *Client {
	c := &Client{
		cfg:    cfg,
		log:    logger,
		client: http.DefaultClient,
	}

	c.log.Debug("api client configuration", zap.Any("cfg", cfg))

	return c
}
