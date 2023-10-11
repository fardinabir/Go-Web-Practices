package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var redisClient *redis.Client

func RedisConnection() *redis.Client {
	if redisClient != nil {
		return redisClient
	}
	_ = InitRedisClient()
	return redisClient
}

func InitRedisClient() error {
	if redisClient != nil {
		return nil
	}
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	//user := viper.GetString("redis.user")
	password := viper.GetString("redis.password")
	//dbname := viper.GetString("redis.dbname")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	err := redisClient.Ping().Err()
	if err != nil {
		redisClient = nil
		return err
	}
	fmt.Println("Redis connection successful...")
	return nil
}

func CheckRateLimit(key string, limit int, duration time.Duration) (bool, error) {
	if redisClient == nil { // If redis is not found then disable rate limiting
		return true, nil
	}
	redisClient.SetNX(key, 1, duration)
	cnt, _ := redisClient.Get(key).Int()
	if cnt <= limit {
		redisClient.Incr(key)
		return true, nil
	}
	return false, nil
}

func CheckRateLimitSliding(key string, limit int, duration time.Duration) (bool, error) {
	now := time.Now()
	item := uuid.New()

	client := redisClient
	minimum := now.Add(-duration)

	// we then remove all requests that have already expired on this set
	client.ZRemRangeByScore(key, "0", strconv.FormatInt(minimum.UnixMilli(), 10))

	// count how many non-expired requests we have on the sorted set
	count := client.ZCount(key, "-inf", "+inf")
	totalRequests, _ := count.Result()
	if int(totalRequests) >= limit {
		return false, nil
	}

	// we add the current request
	client.ZAdd(key, redis.Z{
		Score:  float64(now.UnixMilli()),
		Member: item.String(),
	})

	return true, nil
}
