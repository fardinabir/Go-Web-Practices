package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func RedisConnection() (*redis.Client, error) {
	if RedisClient != nil {
		return RedisClient, nil
	}

	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	//user := viper.GetString("redis.user")
	password := viper.GetString("redis.password")
	//dbname := viper.GetString("redis.dbname")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	err := RedisClient.Ping().Err()
	if err != nil {
		RedisClient = nil
		fmt.Println("Redis connection err : ", err.Error())
		return nil, err
	}
	fmt.Println("Redis connection successful...")
	return RedisClient, nil
}
