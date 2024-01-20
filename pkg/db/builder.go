package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type QBuilder struct {
	pool *pgxpool.Pool
}

func NewQBuilder(log *zap.Logger, cfg *Config) (*QBuilder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if cfg.AutoCreate {
		exists, err := CreateDatabase(ctx, cfg.URL)
		if err != nil {
			log.Error("failed to create database", zap.Error(err))

			return nil, err
		}
		if exists {
			log.Info("the database already exists")
		} else {
			log.Info("the database was created successfully")
		}
	}

	psqlCfg, err := pgx.ParseConfigWithOptions(cfg.URL, pgx.ParseConfigOptions{})
	if err != nil {
		log.Error("failed to parse postgres config", zap.Error(err))

		return nil, err
	}

	conn, err := pgxpool.New(ctx, psqlCfg.ConnString())
	if err != nil {
		log.Error("pool constructor error", zap.Error(err))

		return nil, err
	}

	return &QBuilder{conn}, nil
}

func (qb QBuilder) Pool() *pgxpool.Pool {
	return qb.pool
}
