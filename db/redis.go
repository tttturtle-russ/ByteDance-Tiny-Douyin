package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var redisClient *redis.Client
var ctx context.Context

func initRedis() {
	ctx = context.Background()
	username := viper.GetString("database.redis.username")
	password := viper.GetString("database.redis.password")
	addr := viper.GetString("database.redis.addr")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
	})
}

// Get 获取一个key
func Get(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}

// Set 设置一个key
func Set(key string, value interface{}, timeout time.Duration) error {
	return redisClient.Set(ctx, key, value, timeout).Err()
}

// Del 删除一个key
func Del(key string) error {
	return redisClient.Del(ctx, key).Err()
}

// Exists 返回一个key是否存在，存在返回true，不存在返回false
func Exists(key string) (bool, error) {
	n, err := redisClient.Exists(ctx, key).Result()
	if n == 0 {
		return false, err
	}
	return true, err
}
