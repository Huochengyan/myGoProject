package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"myProject/corll"
	"os"
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
	router.Static("views", "views")
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
		v1.POST("/updateuser", corll.Updateuser)
		v1.POST("/deluser", corll.Deluser)
	}
	v2 := router.Group("/file")
	{
		v2.POST("/uploadfile", corll.Uploadfile)
		v2.GET("/downloadfile", corll.Downloadfile)

	}

	router.Run(":8888") // listen and serve on 0.0.0.0:8080
	fmt.Println("strart ......")
}
