package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func main() {
	fmt.Println("start.................")
	redisopt := redis.Options{
		Addr:     "127.0.0.1:19000",
		DB:       0,
		Password: "123456",
	}

	var redisClient = redis.NewClient(&redisopt)

	commandResult := redisClient.FlushDB()
	fmt.Println(commandResult)
	fmt.Println("=============================")

	for i := 0; i < 1000; i++ {
		keyName := "key" + strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(i)
		fmt.Println(keyName)
		resultcmd := redisClient.Set(keyName, i, time.Minute*9)
		fmt.Println(resultcmd.String())
	}

	//time.Sleep(time.Second*5)
	//
	//resultCmd:=redisClient.FlushDB()
	//fmt.Println(resultCmd)

}
