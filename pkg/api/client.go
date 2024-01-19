package api

import (
	"go.uber.org/zap"
	"net/http"
)

type Client struct {
	cfg    *Config
	log    *zap.Logger
	client *http.Client
}

func NewClient(cfg *Config, logger *zap.Logger) *Client {
	return &Client{
		cfg:    cfg,
		log:    logger,
		client: http.DefaultClient,
	}
}
