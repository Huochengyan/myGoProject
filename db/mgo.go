package db

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

/* mongodb */
func InitMongoDB() (collection *mongo.Database, err error) {
	if Mgo == nil {
		Mgo = connectDB()
	}
	return Mgo, nil
}

var Mgo *mongo.Database

func connectDB() (collection *mongo.Database) {
	const url = "mongodb://192.168.1.108:27017"
	const dbName = "mycol"
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
	collection = client.Database(dbName)
	return collection
}

/* mongodb table Name */
const (
	User string = "user"
	//BlockTxs string = "blockTxs"
	//BlockData string = "blockData"
)
