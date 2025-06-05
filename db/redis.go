package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisMI struct {
	Client *redis.Client
}

var Redis = RedisMI{}

func ConnectRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Redis.Client = rdb
}

func (r *RedisMI) Ping() error {
	ctx := context.Background()
	pong, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println(pong)
	return nil
}

func (r *RedisMI) Set(key string, value []byte) error {
	ctx := context.Background()
	_, err := r.Client.Set(ctx, key, value, time.Second*5).Result()
	if err != nil {
		return err
	}
	return nil
}
