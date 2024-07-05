package cache

import (
	"fmt"
	"gin-rest-api/config"
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

// 初始化redis
func InitRedis(cfg *config.App) error {
	var (
		host     = cfg.Section("redis").Key("host").String()
		port     = cfg.Section("redis").Key("port").String()
		password = cfg.Section("redis").Key("password").String()
	)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Password:     password,
		DB:           0,
		PoolSize:     500,
		MaxRetries:   2,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		MinIdleConns: 50,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 cache 服务器
	_, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
