package corll

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

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

/* Test */
type Test struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	Gender     int    `json:"gender"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	CreateTime int32  `json:"crateTime"`
	Height     int32  `json:"height"`
}

/* role */
type Role struct {
	Id       primitive.ObjectID `bson:"_id"`
	Rolename string             `json:"rolename"`
}

type LoginInfo struct {
	User  *User
	Token string
}
