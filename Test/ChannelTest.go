package main

import "fmt"

func main() {

	var SCount int
	for i := 0; i < 10; i++ {
		SCount = funB(SCount)
	}
	fmt.Println(SCount)
}
func funB(count int) int {
	count = count + 1
	return count
}
