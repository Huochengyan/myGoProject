package db

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"myProject/myProjectUtils"
	"time"
)

const dbName = "mycol"

/* mongodb */
func InitMongoDB() (collection *mongo.Database) {
	if mgo == nil {
		mgo = connectDB()
	}
	return mgo
}

var mgo *mongo.Database

func connectDB() (collection *mongo.Database) {
	var url = myProjectUtils.GetConf("mongo", "url")
	client, err := mongo.NewClient(url)
	if err != nil {
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
	User string = "user"
	Role string = "role"
)
