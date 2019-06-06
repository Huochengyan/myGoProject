package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

func consumer(cname string, ch chan string) {

	//可以循环 for i := range ch 来不断从 channel 接收值，直到它被关闭。

	for i := range ch {
		fmt.Println("consumer-----------", cname, ":", i)
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("producer--", pname, ":", i)
		var value string
		value = strconv.Itoa(i) + "_值"
		ch <- value
	}
}

//func main() {
//	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
//	ch := make(chan string)
//	go producer("pro1", ch)
//	//go producer("pro2", ch)
//	go consumer("con1", ch)
//	//go consumer("con2", ch)
//	time.Sleep(10 * time.Second)
//	close(ch)
//	time.Sleep(10 * time.Second)
//}

func main() {
	//numStr := "1e+27"
	//decimalNum, err := decimal.NewFromString(numStr)
	//if err != nil {
	//	log.Errorf("decimal.NewFromString error, numStr:%s, err:%v", numStr, err)
	//	return
	//}
	//fmt.Println(decimalNum.String())
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromFloat(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

	de1, err := decimal.NewFromString("1e+8")

	de2, err := decimal.NewFromString("1e+8")

	if err == nil {
		fmt.Println("和：" + de1.Add(de2).String())
	}

}
