package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"myProject/controllers"
	"myProject/db"
	"myProject/middleware/jwt"
	"myProject/routers/api"
	"os"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default() //gin.New() //
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.Use(middleware())

	//log -------------------------------------------
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//-----------------------------------------------------------
	//渲染html页面
	// 静态资源加载，本例为css,js以及资源图片
	router.Static("/myproject", "views/static")
	//router.Static("/myproject", "view/")
	//router.Static("static/*")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/* 获得授权Token */
	router.GET("/auth", api.GetAuth)

	t1 := controllers.UserC{Mgo: db.InitMongoDB2()}
	r1 := controllers.RoleC{Mgo: db.InitMongoDB2()}
	//use mongo db
	v1 := router.Group("/user/")
	{
		v1.POST("/login", t1.Login)
		v1.Use(jwt.JWT())
		{
			v1.POST("/insertuser", t1.Insertuser)
			v1.POST("/queryalluser", t1.Queryalluser)
			v1.GET("/getalluser", t1.Getalluser)
			v1.GET("/QueryByUsername", t1.QueryByUsername)
			v1.POST("/updateuser", t1.Updateuser)
			v1.POST("/deluser", t1.Deluser)
			v1.GET("/getroles", r1.GetRoles)
		}
	}
	v2 := router.Group("/file")
	{
		v2.POST("/uploadfile", controllers.Uploadfile)
		v2.GET("/downloadfile", controllers.Downloadfile)
		v2.GET("/readfile", controllers.Readfile)
	}

	//use mysql db
	v3 := router.Group("/userinfo")
	{
		v3.GET("/getuserinfo", controllers.GetUserInfo)
		v3.POST("/updateuser", controllers.UpdateUser)
	}

	t := controllers.TestC{Mgo: db.InitMongoDB2()}
	test := router.Group("/test/")
	{
		test.GET("/test1", t.TestInsert)
		test.GET("/test2", t.Test2)
		test.GET("/test3", t.Test3)
		test.GET("/test4", t.Test4)
		test.GET("/test5", t.Test5)
		test.GET("/test6", t.Test6)
		test.POST("/test7", t.Test7)
	}

	return
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)
		fmt.Println("postBody:", string(x))
		c.Next()
		return
	}
}
