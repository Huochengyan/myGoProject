package corll

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/* user */
type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Address  string             `json:"address"`
	Gender   int                `json:"gender"`
	Email    string             `json:"email"`
	Phone    string             `json:"phone"`
}
