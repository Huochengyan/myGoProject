package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myGoProjectNew/controllers"
	"myGoProjectNew/db"
	"myGoProjectNew/middleware/jwt"
	"myGoProjectNew/routers/api"
	"net/http"
	"os/exec"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func InitRouter() (router *gin.Engine) {
	router = gin.Default() //gin.New() //
	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//router.Use(cors.New(config))

	//跨域
	router.Use(Cors())
	//拦截所有请求 打印下
	//router.Use(middleware())

	//log -------------------------------------------
	// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()
	//// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//-----------------------------------------------------------
	//渲染html页面
	// 静态资源加载，本例为css,js以及资源图片
	//router.Static("/myproject", "views/static")
	router.Static("/myproject", "views/static")
	fmt.Println("open", "http://localhost:8888/myproject/")
	//router.Static("/myproject", "view/")
	//router.Static("static/*")

	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	/* 获得授权Token */
	router.GET("/v1/auth", api.GetAuth)

	t1 := controllers.UserC{Mgo: db.InitMongoDB2()}
	//r1 := controllers.RoleC{Mgo: db.InitMongoDB2()}
	//use mongo db
	v1 := router.Group("/v1/user/")
	{
		v1.POST("/login", t1.Login)
		v1.Use(jwt.JWT())
		//v1.Use()
		{
			v1.GET("/getalluser", t1.Getalluser)
			//v1.POST("/insertuser", t1.Insertuser)
			//v1.POST("/queryalluser", t1.Queryalluser)
			//
			//v1.GET("/QueryByUsername", t1.QueryByUsername)
			//v1.POST("/updateuser", t1.Updateuser)
			//v1.POST("/deluser", t1.Deluser)
			//v1.GET("/getroles", r1.GetRoles)
			v1.POST("/getUserList", t1.GetUserList)
			v1.POST("/update", t1.Update)
			v1.POST("/delete", t1.DelUser)
			v1.POST("/add", t1.AddUser)
		}

	}

	pcinfo := router.Group("/v1/pcinfo/")
	{
		pcinfo.Use(jwt.JWT())
		{
			pcinfo.POST("getPcResourceList", t1.GetPcResourceList)
		}
	}

	logRouter := router.Group("/v1/log/")
	{
		logRouter.Use(jwt.JWT())
		{
			logRouter.POST("getLoginLogList", t1.GetLoginLogList)
		}
	}
	//v2 := router.Group("/file")
	//{
	//	v2.POST("/uploadfile", controllers.Uploadfile)
	//	v2.GET("/downloadfile", controllers.Downloadfile)
	//	v2.GET("/readfile", controllers.Readfile)
	//}

	//use mysql db
	//v3 := router.Group("/userinfo")
	//{
	//	v3.GET("/getuserinfo", controllers.GetUserInfo)
	//	v3.POST("/updateuser", controllers.UpdateUser)
	//}

	//t := controllers.TestC{Mgo: db.InitMongoDB2()}
	//test := router.Group("/test/")
	//{
	//	//test.GET("/test1", t.TestInsert)
	//	//test.GET("/test2", t.Test2)
	//	//test.GET("/test3", t.Test3)
	//	//test.GET("/test4", t.Test4)
	//	//test.GET("/test5", t.Test5)
	//	////test.GET("/test6", t.Test6)
	//	////test.POST("/test7", t.Test7)
	//	//test.GET("/test8", t.GetTestAddMyql)
	//}

	cmd := exec.Command("explorer", "http://localhost:8888/myproject/")
	err2 := cmd.Start()
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	return
}
