package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/ramamimu/go-everything/2-mockery/repository"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedis_Negative(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "product-1"
	value := map[string]interface{}{
		"product_id": "1",
		"amount":     uint(1),
		"price":      float64(2000),
	}
	mock.ExpectHSet(key, value).SetErr(errors.New("error while set hash"))

	cacher := repository.NewCacher(db)
	err := cacher.AddProduct(context.Background(), "1", 1, 2000)
	assert.Errorf(t, err, "expected error while set HSET and return error %s", err.Error())
	assert.ErrorContains(t, err, "error while set hash")
}

func TestRedis_Positive(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "product-1"
	value := map[string]interface{}{
		"product_id": "1",
		"amount":     uint(1),
		"price":      float64(2000),
	}
	mock.ExpectHSet(key, value).SetVal(int64(1))

	cacher := repository.NewCacher(db)
	err := cacher.AddProduct(context.Background(), "1", 1, 2000)
	assert.NoErrorf(t, err, "expected no error while set HSET but got %v", err)
}

func TestRedis_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	opts := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(&opts)

	var err error
	err = client.Ping(context.Background()).Err()
	assert.NoError(t, err, "expected no error while connect to redis")

	cacher := repository.NewCacher(client)
	err = cacher.AddProduct(context.Background(), "a", 12, 9000)
	assert.NoError(t, err, "expected no error while add product to redis")
}
