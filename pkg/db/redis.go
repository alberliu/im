package db

import (
	"im/config"
	"im/pkg/logger"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var RedisCli *redis.Client

func init() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr: config.LogicConf.RedisIP,
		DB:   0,
	})

	_, err := RedisCli.Ping().Result()
	if err != nil {
		logger.Sugar.Error("redis err ", zap.Error(err))
	}
}

func InitRedis(addr string) {
	RedisCli = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	_, err := RedisCli.Ping().Result()
	if err != nil {
		logger.Sugar.Error("redis err ")
		panic(err)
	}
}
