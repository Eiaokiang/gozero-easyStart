package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func NewRedisStore() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("database.redis.addr"),
		Password: viper.GetString("database.redis.password"),
		DB:       viper.GetInt("database.redis.db"),
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("redis连接失败")
	}
	return client
}
