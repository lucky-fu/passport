// Package container ...
//
// File:  redis
//
// Description: redis
//
// Date: 2021/5/23 4:04 下午
package container

import (
	"fmt"
	"lib/src/github.com/go-redis/redis"
	"time"

	"github.com/passport/define"
)

var (
	RedisClient *redisClient
)

type redisClient struct {
	instance *redis.Client
}

// InitRedis ...
func InitRedis() {
	var (
		err         error
		redisConfig *define.ConfigRedis
	)

	if redisConfig, err = GetRedisConfig("config/redis"); err != nil {
		panic("get redis config is failed")
	}

	RedisClient.instance = redis.NewClient(&redis.Options{
		DB:           redisConfig.DB,
		Addr:         fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		DialTimeout:  time.Duration(redisConfig.ConnectTimeout*1000) * time.Millisecond,
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout*1000) * time.Millisecond,
		WriteTimeout: time.Duration(redisConfig.WriteTimeout*1000) * time.Millisecond,
		Password:     redisConfig.Password,
	})

	if RedisClient.instance == nil {
		panic("redis init is failed")
	}
}

// GetRedisClient ...
func (rc *redisClient) GetRedisClient() *redis.Client {
	return rc.instance
}

// Get ...
func (rc *redisClient) Get(key string) ([]byte, error) {
	return rc.instance.Get(key).Bytes()
}
