package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myProject/myProjectUtils"
	"time"
)

/*
   mongodb访问方法。

*/
const dbName = "mycol"

/* mongodb */
//func InitMongoDB() (collection *mongo.Database) {
//	if mgo == nil {
//		mgo = connectDB()
//	}
//	return mgo
//}
/* mongodb */
func InitMongoDB2() (collection *mongo.Database) {
	if mgo == nil {
		mgo = connectDB()
	}
	return mgo
}

var mgo *mongo.Database

func connectDB() (collection *mongo.Database) {
	//var url = myProjectUtils.GetConf("mongo", "url")
	//opts := options.ClientOptions{
	//	Hosts: []string{url}
	//}

	config := myProjectUtils.GetConfInfo("mongo")

	opts := options.ClientOptions{Hosts: []string{config.Key("url").String()}}
	credential := options.Credential{
		Username: config.Key("username").String(), Password: config.Key("password").String(),
		AuthSource: config.Key("db").String(),
	}
	opts.Auth = &credential

	client, err := mongo.NewClient(&opts)
	if err != nil {
		panic(err)
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil
	}
	var dbName = myProjectUtils.GetConf("mongo", "db")
	collection = client.Database(dbName)
	return collection
}

/* mongodb table Name */
const (
	User   string = "user"
	Role   string = "role"
	Test   string = "test"
	Test1  string = "test1"
	Test2  string = "test2"
	PcInfo string = "pcinfo"
)
