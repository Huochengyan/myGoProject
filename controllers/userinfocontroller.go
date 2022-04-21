package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myGoProjectNew/db"
	"net/http"
)

//获取所有用户
func GetUserInfo(g *gin.Context) {
	rsp := new(Rsp)
	db := db.InitMysqlDb()

	Rows, err := db.Query("select  * from  userinfo;")
	fmt.Println(Rows, err)
	var user = new(UserInfo)
	var users []UserInfo
	for Rows.Next() {
		Rows.Scan(&user.Id, &user.UserName)
		users = append(users, *user)
	}
	rsp.Code = 200
	rsp.Data = users
	g.JSON(http.StatusOK, rsp)
}

func UpdateUser(g *gin.Context) {
	rsp := new(Rsp)
	rsp.Code = 200
	rsp.Data = nil
	g.JSON(http.StatusOK, rsp)
}

type UserInfo struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}
