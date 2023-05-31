package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func RedisConnection() *redis.Client {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	//user := viper.GetString("redis.user")
	password := viper.GetString("redis.password")
	//dbname := viper.GetString("redis.dbname")

	client := redis.NewClient(&redis.Options{
		Addr:     host + port,
		Password: password,
		DB:       0,
	})

	fmt.Println("Redis connection successful...")
	return client
}
