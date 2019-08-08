package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	addrStr := make([]string, 0)
	addrStr = append(addrStr, "127.0.0.1:7000")
	addrStr = append(addrStr, "127.0.0.1:7001")
	addrStr = append(addrStr, "127.0.0.1:7002")
	addrStr = append(addrStr, "127.0.0.1:7003")
	addrStr = append(addrStr, "127.0.0.1:7004")
	addrStr = append(addrStr, "127.0.0.1:7005")

	redisopt := redis.ClusterOptions{
		Addrs:    addrStr,
		Password: "",
	}

	clusterClient := redis.NewClusterClient(&redisopt)

	clusterClient.Set("key1", "123", 100)

	result := clusterClient.Get("key1").String()
	fmt.Println(result)

	//out :123

}
