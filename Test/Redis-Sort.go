package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	redisopt := redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "123456",
	}

	var redisClient = redis.NewClient(&redisopt)
	fmt.Println(redisClient.String())

	//n,err := c.Do("zadd","key","score","member")  //写
	//result,err := redis.Values(c.Do("zrange","key",0,-1))//读

	cmd := redisClient.Do("zadd", "key1", "0", "zhangsan1")

	fmt.Println(cmd)

	cmd = redisClient.Do("zadd", "key1", "8", "zhangsan2")
	fmt.Println(cmd)

	cmd = redisClient.Do("zadd", "key2", "8", "zhangsan")

	fmt.Println(redisClient.Keys(""))
}
