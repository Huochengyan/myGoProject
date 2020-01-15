package main

import (
	"flag"
	"fmt"
	_ "github.com/fvbock/endless"
	"github.com/go-ini/ini"
	"github.com/robfig/cron"
	"io/ioutil"
	"myProject/log"
	"myProject/myProjectUtils"
	"myProject/routers"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	log.Info(time.Now().Format("2006-01-02 15:04:05") + "strart ......")
	flag.Parse()

	//cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, "conf/app.ini")
	//if err != nil {
	//	panic(err)
	//}
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		panic(err)
	}

	myProjectUtils.Config = cfg

	router := routers.InitRouter()
	//cronInit()

	port := cfg.Section("http").Key("port").String()

	err1 := router.Run(port)
	if err1 != nil {
		panic(err1)
	}
	log.Info("strart ......")

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

//定时器
func cronInit() {
	go func() {
		crontab := cron.New()
		crontab.AddFunc("*/20 * * * *", myfunc) //5S
		crontab.Start()
	}()
}

// 加个定时器
func myfunc() {
	//fmt.Println("5秒打印一次！！")
	url := "https://blog.csdn.net/u010919083/article/details/79358775"
	gethtml(url)
	url2 := "https://blog.csdn.net/u010919083/article/details/78210868"
	gethtml(url2)
}
func gethtml(url string) {

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "PostmanRuntime/7.13.0")
	req.Header.Add("Accept", "*/*")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
