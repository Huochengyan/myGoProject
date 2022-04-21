package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/date"
)

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

type PcInfo struct {
	Id            primitive.ObjectID `bson:"_id"`
	GetCpuPercent float32            `json:"getcpupercent"`
	DiskPercent   float32            `json:"diskpercent"`
	MemPercent    float32            `json:"mempercent"`
	CreateTime    date.Date          `json:"crateTime"`
}
