package main

import (
	"fmt"
	"myProject/routers"
	"os/exec"
	"runtime"
)

func main() {
	router := routers.InitRouter()
	err := router.Run(":8888") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
	fmt.Println("strart ......")

	//调用浏览器打开页面
	//OpenUrl("https://www.baidu.com")
	//browser.OpenURL("http://www.baidu.com")
}

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenUrl(uri string) {
	run, _ := commands[runtime.GOOS]
	exec.Command(run, uri).Start()
}
