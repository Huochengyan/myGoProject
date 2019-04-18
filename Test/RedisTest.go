package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"thkAdmin/utils"
	"time"
)

func main() {
	/* 获取  token 后存起来*/
	redisopt := redis.Options{
		Addr:     utils.GetConf("redis", "url"),
		DB:       0,
		Password: "",
	}

	var redisClient = redis.NewClient(&redisopt)
	redisClient.PubSubChannels("123")

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 1)
		redisClient.Publish("123", i)
		fmt.Println(i)
	}

}
