package initialize

import (
	"github.com/redis/go-redis/v9"
	"serve/global"
)

func InitRedis1() {
	client := redis.NewClient(&redis.Options{
		//Addr:     "127.0.0.1:6377",
		Addr:     "154.83.15.174:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	global.Backend_REDIS = client

}
