package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"myProject/corll"
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

	t1 := corll.UserC{Mgo: db.InitMongoDB2()}
	r1 := corll.RoleC{Mgo: db.InitMongoDB2()}
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

	t := corll.TestC{Mgo: db.InitMongoDB2()}
	test := router.Group("/test/")
	{
		test.GET("/test1", t.TestInsert)
		test.GET("/test2", t.Test2)
		test.GET("/test3", t.Test3)
	}
	return
}
