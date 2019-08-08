package main

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"myProject/routers"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	router := routers.InitRouter()
	cronInit()

	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		panic(err)
	}

	err1 := router.Run(cfg.Section("http").Key("port").String()) // listen and serve on 0.0.0.0:8080
	if err1 != nil {
		panic(err1)
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

	//req.Header.Add("User-Agent", "PostmanRuntime/7.13.0")
	//req.Header.Add("Accept", "*/*")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body)[:20])
}
