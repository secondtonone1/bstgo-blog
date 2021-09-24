package redis

import (
	"bstgo-blog/config"
	"log"

	"github.com/go-redis/redis"
)

var rediscli *redis.Client

func InitRedis() {
	rediscli = redis.NewClient(&redis.Options{
		Addr:         config.TotalCfgData.Redis.Host,
		Password:     config.TotalCfgData.Redis.Passwd,
		DB:           config.TotalCfgData.Redis.DB,
		PoolSize:     config.TotalCfgData.Redis.PoolSize,
		MinIdleConns: config.TotalCfgData.Redis.IdleCons,
	})

	_, err := rediscli.Ping().Result()
	if err != nil {
		log.Println("ping failed, error is ", err)
		return
	}

	log.Println("redis init success!!!")
}
