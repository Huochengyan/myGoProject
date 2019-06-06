package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string, 100)
	//ch<- "123"
	//ch<- "456"
	//
	//bo:=<-ch
	//fmt.Println(bo)
	go pro(ch)
	go cus(ch)
}

func pro(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- strconv.Itoa(i)
	}
}
func cus(ch chan string) {
	for {
		fmt.Println(ch)
	}
}
