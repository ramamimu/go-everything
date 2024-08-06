package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// redis.Cmdable is interface from redis
type Cache struct {
	cacher redis.Cmdable
}

func NewCacher(c redis.Cmdable) *Cache {
	return &Cache{
		cacher: c,
	}
}

func (c Cache) AddProduct(ctx context.Context, productID string, amount uint, price float64) error {
	key := fmt.Sprintf("product-%s", productID)
	value := map[string]interface{}{
		"product_id": productID,
		"amount":     amount,
		"price":      price,
	}

	err := c.cacher.HSet(ctx, key, value).Err()

	if err != nil {
		return err
	}

	return nil
}
