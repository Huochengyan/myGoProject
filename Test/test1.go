package main

import (
	"fmt"
	"time"
)

func consumer(cname string, ch chan int) {

	//可以循环 for i := range ch 来不断从 channel 接收值，直到它被关闭。

	for i := range ch {
		fmt.Println("consumer-----------", cname, ":", i)
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println("producer--", pname, ":", i)
		ch <- i
	}
}

func main() {
	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
	ch := make(chan int)
	go producer("生产者1", ch)
	go producer("生产者2", ch)
	go consumer("消费者1", ch)
	go consumer("消费者2", ch)
	time.Sleep(10 * time.Second)
	close(ch)
	time.Sleep(10 * time.Second)
}
