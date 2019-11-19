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
