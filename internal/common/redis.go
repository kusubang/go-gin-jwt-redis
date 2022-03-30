package common

import (
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var client *redis.Client

func init() {
	dsn := viper.GetString("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return client
}
