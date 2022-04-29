package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/robfig/cron"
	"myGoProjectNew/myProjectUtils"
	"myGoProjectNew/routers"
	"myGoProjectNew/utils"
	"os/exec"
	"runtime"
	"time"
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
	cronInit()

	port := cfg.Section("http").Key("port").String()
	err1 := router.Run(port)
	if err1 != nil {
		fmt.Println(err)
	}

	urlLocal := "http://localhost:" + port + "/myproject/"
	fmt.Println("open", urlLocal)
	cmd := exec.Command("explorer", urlLocal)
	err2 := cmd.Start()
	if err2 != nil {
		fmt.Println(err2.Error())
	}
}

//定时器
func cronInit() {
	go func() {
		crontab := cron.New()
		crontab.AddFunc("* */10 * * *", myfunc) //1 分钟
		crontab.Start()
	}()
}

// 加个定时器
func myfunc() {
	fmt.Println(time.Now(), "10秒打印一次！！")

	fmt.Println(runtime.NumCPU())
	fmt.Println(utils.GetCpuPercent())
	fmt.Println(utils.GetDiskPercent())
	fmt.Println(utils.GetMemPercent())

}
