package main

import (
	"context"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myGoProjectNew/db"
	"myGoProjectNew/models"
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
		log.Error(err)
	}

	urlLocal := "http://localhost" + port + "/myproject/"
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
		crontab.AddFunc("*/5 * * * *", myfunc) //1 分钟
		crontab.Start()
	}()
}

// 加个定时器
func myfunc() {
	fmt.Println(time.Now(), "10秒打印一次！！")

	fmt.Println("CPU Number:", runtime.NumCPU())
	fmt.Println("CPU Use:", utils.GetCpuPercent(), "%")
	fmt.Println("Disk Use:", utils.GetDiskPercent(), "%")
	fmt.Println("Memory user:", utils.GetMemPercent(), "%")

	info := models.PcResource{Id: primitive.NewObjectID(),
		DiskPercent: utils.GetDiskPercent(), CpuPercent: utils.GetCpuPercent(), CreateTime: time.Now().Unix()}
	result, err := db.InitMongoDB2().Collection(db.PcResource).InsertOne(context.Background(), info)
	if err == nil {
		fmt.Println(result.InsertedID)
	}
}
