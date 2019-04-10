package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"myProject/corll"
	"os"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default() //gin.New() //
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

	//use mysql db
	v3 := router.Group("/userinfo")
	{
		v3.GET("/getuserinfo", corll.GetUserInfo)
		v3.POST("/updateuser", corll.UpdateUser)
	}
	return
}
