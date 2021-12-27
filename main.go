package main

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/robfig/cron"
	"myProject/log"
	"myProject/myProjectUtils"
	"myProject/routers"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	log.Info(time.Now().Format("2006-01-02 15:04:05") + "strart ......")
	flag.Parse()

	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		panic(err)
	}

	myProjectUtils.Config = cfg

	router := routers.InitRouter()
	//定时器应用
	cronInit()

	port := cfg.Section("http").Key("port").String()
	err1 := router.Run(port)
	if err1 != nil {
		panic(err1)
	}
	log.Info("strart ......")
	fmt.Println("open", "http://127.0.0.1:8888/myproject/")

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

//定时器
func cronInit() {
	go func() {
		crontab := cron.New()
		crontab.AddFunc("*/1 * * * *", myfunc) //1 分钟
		crontab.Start()
	}()
}

// 加个定时器
func myfunc() {
	fmt.Println(time.Now(), "5秒打印一次！！")

	//fmt.Println(runtime.NumCPU())
	//fmt.Println(utils.GetCpuPercent())
	//fmt.Println(utils.GetDiskPercent())
	//fmt.Println(utils.GetMemPercent())
}
