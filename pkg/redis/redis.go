package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis/recipe"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Recipe *recipe.Cache

	// DurCacheUpdate duration in minutes to update cache
	durCacheUpdate int
	db             *postgres.Postgres
}

func NewRedis(ctx context.Context, cfg config.RedisConfig, db *postgres.Postgres) (*Redis, error) {
	r := &Redis{
		durCacheUpdate: cfg.DurCacheUpdate,
		db:             db,
	}

	c := redis.NewClient(&redis.Options{
		Addr:            cfg.Addr,
		Password:        cfg.Password,
		DB:              cfg.DB,
		DialTimeout:     time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:     time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(cfg.WriteTimeout) * time.Second,
		PoolSize:        cfg.PoolSize,
		MinIdleConns:    cfg.MinIdleConns,
		PoolTimeout:     time.Duration(cfg.PoolTimeout) * time.Second,
		TLSConfig:       nil,
		MaxRetries:      cfg.MaxRetries,
		MinRetryBackoff: time.Duration(cfg.MinRetryBackoff) * time.Millisecond,
		MaxRetryBackoff: time.Duration(cfg.MaxRetryBackoff) * time.Millisecond,
	})

	_, err := c.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	r.Recipe = recipe.NewRecipe(c)

	return r, nil
}
