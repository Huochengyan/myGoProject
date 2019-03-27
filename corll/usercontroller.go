package corll

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"myProject/db"
	"net/http"
	"strconv"
)

/*login*/
func Login(g *gin.Context) {
	fmt.Println("login.........")
	rsp := new(Rsp)
	name := g.PostForm("username")
	pass := g.PostForm("password")
	mgo, err := db.InitMongoDB()
	findfilter := bson.D{{"username", g.PostForm("username")}, {"password", g.PostForm("password")}}
	cur, err := mgo.Collection(db.User).Find(context.Background(), findfilter)
	for cur.Next(context.Background()) {
		elme := new(User)
		err := cur.Decode(elme)
		if err == nil {
			if elme.Username == name && elme.Password == pass {
				rsp.Msg = "success"
				rsp.Code = 200
				g.JSON(http.StatusOK, rsp)
				return
			}
		}
	}

	rsp.Msg = "user is null"
	rsp.Code = 201
	rsp.Data = err
	g.JSON(http.StatusOK, rsp)
}

/* insert user table */
func Insertuser(g *gin.Context) {
	rsp := new(Rsp)
	fmt.Println("InsertUser.........")
	if g.PostForm("username") == "" || g.PostForm("password") == "" {
		rsp.Msg = "user is null"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
	}
	mgo, err := db.InitMongoDB()
	if err != nil {
		fmt.Println(nil)
	}
	newuser := new(User)
	newuser.Username = g.PostForm("username")
	newuser.Password = g.PostForm("password")
	insertID, err := mgo.Collection(db.User).InsertOne(context.Background(), newuser)
	fmt.Println(insertID)
	if err == nil {
		rsp.Msg = "success"
		rsp.Code = 200
		g.JSON(http.StatusOK, rsp)
		return
	} else {
		rsp.Msg = "faild"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
}

/* query all user */
func Queryalluser(g *gin.Context) {
	fmt.Println("Queryalluser.........")
	rsp := new(Rsp)
	mgo, err := db.InitMongoDB()
	if err != nil {
		fmt.Println("connect db err:", err)
	}
	var users []User
	cur, err := mgo.Collection(db.User).Find(context.Background(), bson.D{}, nil)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(User)
			err := cur.Decode(elme)
			if err == nil {
				users = append(users, *elme)
			}
		}
	}
	rsp.Msg = "success"
	rsp.Code = 200
	rsp.Data = users
	g.JSON(http.StatusOK, rsp)
	return
}

/* get all user */
func Getalluser(g *gin.Context) {
	fmt.Println("Getalluser.........")
	rsp := new(Rsp)
	mgo, err := db.InitMongoDB()
	if err != nil {
		fmt.Println("connect db err:", err)
	}
	filter := bson.M{}

	int, err := strconv.Atoi(g.PostForm("gender"))
	if err == nil {
		filter = bson.M{"gender": int}
	}
	username := g.PostForm("username")
	if username != "" {
		filter = bson.M{"username": username}
	}

	var users []User
	cur, err := mgo.Collection(db.User).Find(context.Background(), filter, nil)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(User)
			err := cur.Decode(elme)
			if err == nil {
				users = append(users, *elme)
			}
		}
	}

	rsp.Msg = "success"
	rsp.Code = 0
	rsp.Data = users
	g.JSON(http.StatusOK, rsp)
	return
}

/* update user */
func Updateuser(g *gin.Context) {
	fmt.Println("update user.................")
	rsp := new(Rsp)
	mgo, err := db.InitMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	id := g.PostForm("id")
	oldId, err := primitive.ObjectIDFromHex(id)
	newUsername := g.PostForm("username")

	info, err := mgo.Collection(db.User).UpdateOne(context.Background(), bson.M{"_id": bsonx.ObjectID(oldId)},
		bson.M{"$set": bson.M{"username": newUsername}})
	if err == nil {
		fmt.Println(info.MatchedCount)
	}
	rsp.Msg = "success"
	rsp.Code = 200
	g.JSON(http.StatusOK, rsp)
	return
}

/* deluser*/
func Deluser(g *gin.Context) {
	fmt.Println("update user.................")
	rsp := new(Rsp)
	mgo, err := db.InitMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	username := g.PostForm("username")
	info, err := mgo.Collection(db.User).DeleteOne(context.Background(), bson.M{"username": username})

	if info.DeletedCount == 0 {
		rsp.Msg = "faild"
		rsp.Data = "username is  not exist!!"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}

	if err == nil {
		fmt.Println(info.DeletedCount)
		rsp.Msg = "success"
		rsp.Code = 200
		g.JSON(http.StatusOK, rsp)
		return
	} else {
		rsp.Msg = "faild"
		rsp.Data = err
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
}
