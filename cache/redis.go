package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type RedisDB struct {
	client *redis.Client
}

var ctx = context.Background()

func Connect(host string, port int) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "",
		DB:       0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisDB{
		client: client,
	}, nil
}

func (client *RedisDB) Set( key string, value string, expiration time.Duration)(string, error){
	return client.client.Set(ctx, key,value, expiration).Result()
}

func (client *RedisDB) Get(key string) (string, error){
	return client.client.Get(ctx, key).Result()
}

func (client *RedisDB) Del(key string) (int64, error){
	return client.client.Del(ctx, key).Result()
}

func (client *RedisDB) Exists(key string) (int64, error){
	return client.client.Exists(ctx, key).Result()
}

