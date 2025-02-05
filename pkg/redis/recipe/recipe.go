package recipe

import (
	"context"
	"errors"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
	"github.com/redis/go-redis/v9"
	"github.com/vmihailenco/msgpack/v5"
)

type Cache struct {
	client *redis.Client
	prefix string
}

func NewRecipe(client *redis.Client) *Cache {
	return &Cache{
		client: client,
		prefix: "recipe:",
	}
}

// Set create/update value by key.
func (c *Cache) Set(
	ctx context.Context,
	key string,
	val recipe.InCache,
	expiration time.Duration,
) error {
	b, err := msgpack.Marshal(val)
	if err != nil {
		return err
	}

	p := c.client.Pipeline()
	p.Set(ctx, c.prefix+key, b, expiration)
	if _, err = p.Exec(ctx); err != nil {
		return err
	}

	return nil
}

// Get value by key.
func (c *Cache) Get(ctx context.Context, key string) (recipe.InCache, error) {
	b, err := c.client.Get(ctx, c.prefix+key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return recipe.InCache{}, apperrors.ErrCacheKeyNotExists
		}
		return recipe.InCache{}, err
	}

	var ric recipe.InCache
	if err := msgpack.Unmarshal(b, &ric); err != nil {
		return recipe.InCache{}, err
	}

	return ric, nil
}

// Del key/keys.
func (c *Cache) Del(ctx context.Context, keys ...string) error {
	p := c.client.Pipeline()

	for _, key := range keys {
		p.Del(ctx, c.prefix+key)
	}

	_, err := p.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
