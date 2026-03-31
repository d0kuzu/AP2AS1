package database

import (
	"assign1/internal/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPool(ctx context.Context, cfg *config.Settings) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(cfg.GetDSN())
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
