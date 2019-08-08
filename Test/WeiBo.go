package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	for i := 0; i < 10; i++ {
		url := "https://blog.csdn.net/u010919083/article/details/79358775"
		gethtml(url)
	}

}
func gethtml(url string) {

	req, _ := http.NewRequest("GET", url, nil)

	//req.Header.Add("User-Agent", "PostmanRuntime/7.13.0")
	//req.Header.Add("Accept", "*/*")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body)[:20])
}
