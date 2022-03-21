package main

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"myGoProjectNew/myProjectUtils"
	"myGoProjectNew/routers"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "strart ......")
	flag.Parse()

	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		panic(err)
	}

	myProjectUtils.Config = cfg

	router := routers.InitRouter()
	//定时器应用
	//cronInit()

	port := cfg.Section("http").Key("port").String()
	err1 := router.Run(port)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("strart ......")
	fmt.Println("open", "http://127.0.0.1:8888/myproject/")

}
