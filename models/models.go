package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/* user */
type User struct {
	Id         primitive.ObjectID `bson:"_id"`
	Username   string             `json:"username"`
	Password   string             `json:"password"`
	Address    string             `json:"address"`
	Gender     int                `json:"gender"`
	Email      string             `json:"email"`
	Phone      string             `json:"phone"`
	CreateTime int32              `json:"crateTime"`
}

type PageHelper struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type UserList struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type PcResource struct {
	Id primitive.ObjectID `bson:"_id"`
	//硬盘使用 百分比
	DiskPercent float64 `json:"diskPercent"`
	//CPU使用 百分比
	CpuPercent float64 `json:"cpuPercent"`
	CreateTime int64   `json:"crateTime"`
}
