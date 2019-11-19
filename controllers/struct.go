package controllers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myProject/models"
)

type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/* Test */
type Test struct {
	Id         primitive.ObjectID `bson:"_id"`
	Username   string             `json:"username"`
	Password   string             `json:"password"`
	Address    string             `json:"address"`
	Gender     int                `json:"gender"`
	Email      string             `json:"email"`
	Phone      string             `json:"phone"`
	CreateTime int32              `json:"crateTime"`
	Height     int32              `json:"height"`
}

/* role */
type Role struct {
	Id       primitive.ObjectID `bson:"_id"`
	Rolename string             `json:"rolename"`
}

type LoginInfo struct {
	User  *models.User
	Token string
}

type PostParamTx struct {
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`

	//条件字段 （int 类型字段之所以用String，是应为int字段没有时，或默认为0的）
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
}

type ResultData struct {
	PageNum  int           `json:"pageNum"`
	PageSize int           `json:"pageSize"`
	Pages    int           `json:"pages"`
	Total    int           `json:"total"`
	DataList []models.User `json:"dataList"`
}
