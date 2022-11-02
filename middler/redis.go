package middler

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var Redis *redisCache

type redisCache struct {
	instance *redis.Client
}

// InitRedis ...
func InitRedis() {
	Redis = &redisCache{instance: redis.NewClient(
		&redis.Options{
			DB:           Config.Redis.DB,
			Addr:         fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port),
			DialTimeout:  time.Duration(Config.Redis.Timeout.Connect*1000) * time.Millisecond,
			ReadTimeout:  time.Duration(Config.Redis.Timeout.Read*1000) * time.Millisecond,
			WriteTimeout: time.Duration(Config.Redis.Timeout.Write*1000) * time.Millisecond,
			Password:     Config.Redis.Password,
		})}

	if Redis.instance == nil {
		panic("初始化redis失败")
	}

	fmt.Println("初始化redis成功")
}

// Set ...
func (rc *redisCache) Set(key, value string, ttl time.Duration) error {
	return rc.instance.Set(key, value, ttl).Err()
}

// Get ...
func (rc *redisCache) Get(key string) ([]byte, error) {
	return rc.instance.Get(key).Bytes()
}
