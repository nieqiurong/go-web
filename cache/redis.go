package cache

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"go-web/setting"
	"log"
)

var client *redis.Client

func InitRedis() {
	redisConfig := setting.Application.Redis
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("connect redis fail :", err)
	}
}
