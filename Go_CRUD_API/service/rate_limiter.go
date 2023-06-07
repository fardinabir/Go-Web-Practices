package service

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//func RateLimitPass(userID string, maxRequests int, duration time.Duration, client *redis.Client) bool {
//	key := "ratelimit:" + userID
//	now := time.Now().Unix()
//
//	pipe := client.TxPipeline()
//	pipe.ZRemRangeByScore(key, "0", strconv.Itoa(int(time.Now().Unix())))
//	pipe.ZCard(key)
//	pipe.ZAdd(key, redis.Z{Score: float64(now), Member: now})
//	pipe.Expire(key, duration)
//
//	_, err := pipe.Exec()
//	if err != nil {
//		// Handle error
//		return false
//	}
//
//	result, err := client.ZCard(key).Result()
//	if err != nil {
//		// Handle error
//		return false
//	}
//	fmt.Println(result, "resulttttttttttttttttttttttttt")
//	return result <= int64(maxRequests)
//}

func getTime() (map[string]interface{}, error) {
	var maxLimit int64
	var limitTime int64
	maxLimit = 5
	limitTime = 60
	systemUser := "john_doe"
	uniqueKey := systemUser
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	var counter int64
	counter, err := redisClient.Get(uniqueKey).Int64()
	if err == redis.Nil {
		err = redisClient.Set(uniqueKey, 1, time.Duration(limitTime)*time.Second).Err()
		if err != nil {
			return nil, err
		}
		counter = 1
	} else if err != nil {
		return nil, err
	} else {
		if counter >= maxLimit {
			return nil, errors.New("Limit reached.")
		}
		counter, err = redisClient.Incr(uniqueKey).Result()
		if err != nil {
			return nil, err
		}
	}
	dt := time.Now()
	res := map[string]interface{}{
		"data": dt.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

func CheckLimit(key string, limit int, duration time.Duration, client *redis.Client) (bool, error) {
	p := client.Pipeline()
	incrResult := p.Incr(key)
	ttlResult := p.TTL(key)

	if _, err := p.Exec(); err != nil {
		fmt.Println(err, "failed to execute increment to key %v", key)
		return false, err
	}
	totalRequests, err := incrResult.Result()
	if err != nil {
		fmt.Println(err, "failed to increment key %v", key)
		return false, err
	}
	var ttlDuration time.Duration
	d, err := ttlResult.Result()
	if err != nil || d == (-1*time.Second) {
		ttlDuration = duration
		if err := client.Expire(key, duration).Err(); err != nil {
			fmt.Println(err, "failed to set an expiration to key %v", key)
			return false, err
		}
	} else {
		ttlDuration = d
	}
	//expiresAt := time.Now().Add(ttlDuration)
	requests := uint64(totalRequests)
	if requests > uint64(limit) {
		fmt.Println("Limit will be expired in : ", ttlDuration)
		return false, nil
	}
	return true, nil
}

func RateLimitCheckShort(key string, limit int, duration time.Duration, client *redis.Client) (bool, error) {
	res, _ := client.Get(key).Int64()
	if res > int64(limit) {
		return false, nil
	}
	incrResult := client.Incr(key)
	totalRequests, _ := incrResult.Result()
	if totalRequests == 1 {
		if err := client.Expire(key, duration).Err(); err != nil {
			fmt.Println(err, "failed to set an expiration to key %v", key)
			return false, err
		}
	}
	if totalRequests > int64(limit) {
		return false, nil
	}
	return true, nil
}
