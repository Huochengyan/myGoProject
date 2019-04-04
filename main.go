package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/pkg/browser"
	"io"
	"myProject/corll"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	router := gin.New() //gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	//log -------------------------------------------
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//-----------------------------------------------------------
	//渲染html页面
	// 静态资源加载，本例为css,js以及资源图片
	router.Static("static", "static")

	//router.Static("static/*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/user/")
	{
		v1.POST("/login", corll.Login)
		v1.POST("/insertuser", corll.Insertuser)
		v1.POST("/queryalluser", corll.Queryalluser)
		v1.GET("/getalluser", corll.Getalluser)
		v1.GET("/QueryByUsername", corll.QueryByUsername)
		v1.POST("/updateuser", corll.Updateuser)
		v1.POST("/deluser", corll.Deluser)
	}
	v2 := router.Group("/file")
	{
		v2.POST("/uploadfile", corll.Uploadfile)
		v2.GET("/downloadfile", corll.Downloadfile)
		v2.GET("/readfile", corll.Readfile)

	}

	v3 := router.Group("/userinfo")
	{
		v3.GET("/getuserinfo", corll.GetUserInfo)

	}

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
