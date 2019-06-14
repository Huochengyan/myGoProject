package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	for i := 0; i < 10; i++ {
		url := "http://blog.sina.com.cn/s/blog_ec1a30820102vnc2.html"
		gethtml(url)
	}

}
func gethtml(url string) {

	req, _ := http.NewRequest("GET", url, nil)

	//req.Header.Add("User-Agent", "PostmanRuntime/7.13.0")
	//req.Header.Add("Accept", "*/*")
	//req.Header.Add("Cache-Control", "no-cache")
	//req.Header.Add("Postman-Token", "5e68c2f5-b9a8-4708-b0a4-f75d012597cb,919e97cf-ea51-4f74-8751-c2d6112dd2ab")
	//req.Header.Add("Host", "blog.sina.com.cn")
	//req.Header.Add("accept-encoding", "gzip, deflate")
	//req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
