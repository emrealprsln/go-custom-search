package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func InitCache() (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT")),
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Println("redis successfully configured")

	return rdb, nil
}

func GetValue(key string) string {
	result, _ := rdb.Get(ctx, key).Result()
	return result
}

func SetValue(key string, values interface{}) {
	encoded, _ := json.Marshal(values)
	rdb.Set(ctx, key, encoded, 0)
}
