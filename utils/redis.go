package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedisClient(host string, port int) (*redis.Client, error) {
	var ctx = context.Background()
	if RedisClient == nil {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		_, err := RedisClient.Ping(ctx).Result()
		if err != nil {
			return nil, err
		}
	}
	return RedisClient, nil
}
