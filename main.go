package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"myGoProjectNew/myProjectUtils"
	"myGoProjectNew/routers"
)

func main() {
	//now:=time.Now()
	//fmt.Println(now.Format("2006-01-02 15:04:05") + "strart ......")
	//flag.Parse()

	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println(err)
	}

	myProjectUtils.Config = cfg

	router := routers.InitRouter()
	//定时器应用
	//cronInit()

	port := cfg.Section("http").Key("port").String()
	err1 := router.Run(port)
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println("strart ......")
	fmt.Println("open", "http://127.0.0.1:8888/myproject/")

}
