package corll

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/g/util/gvalid"
	_ "github.com/gogf/gf/g/util/gvalid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
	"myProject/db"
	"myProject/pkg/util"
	"net/http"
	"strconv"
)

/*login user*/
func Login(g *gin.Context) {
	fmt.Println("login.........")
	rsp := new(Rsp)
	name := g.PostForm("username")
	pass := g.PostForm("password")

	var gerr *gvalid.Error
	gerr = gvalid.Check(g.PostForm("username"), "required", nil)
	if gerr != nil {
		rsp.Msg = "faild"
		rsp.Code = 201
		rsp.Data = gerr.Maps()
		g.JSON(http.StatusOK, rsp)
		return
	}
	gerr = gvalid.Check(g.PostForm("password"), "required", nil)
	if gerr != nil {
		rsp.Msg = "faild"
		rsp.Code = 201
		rsp.Data = gerr.Maps()
		g.JSON(http.StatusOK, rsp)
		return
	}

	mgo := db.InitMongoDB()
	findfilter := bson.D{{"username", g.PostForm("username")}, {"password", g.PostForm("password")}}
	cur, err := mgo.Collection(db.User).Find(context.Background(), findfilter)
	for cur.Next(context.Background()) {
		elme := new(User)
		err := cur.Decode(elme)
		if err == nil {
			if elme.Username == name && elme.Password == pass {
				var info = new(LoginInfo)
				info.User = elme
				token, err := util.GenerateToken(g.PostForm("username"), g.PostForm("password"))
				if err == nil {
					info.Token = token
				}
				rsp.Msg = "success"
				rsp.Code = 200
				rsp.Data = info
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
	newuser := new(User)
	var err1 *gvalid.Error
	//1. 单参数校验
	err1 = gvalid.Check(g.PostForm("gender"), "required|integer|between:0,150", nil)
	if err1 != nil {
		rsp.Msg = err1.String()
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
	err1 = gvalid.Check(g.PostForm("username"), "required", nil)
	if err1 != nil {
		rsp.Msg = "username:" + err1.String()
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
	err1 = gvalid.Check(g.PostForm("password"), "required|password", nil)
	if err1 != nil {
		rsp.Msg = "password:" + err1.String()
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
	err1 = gvalid.Check(g.PostForm("email"), "required|email", nil)
	if err1 != nil {
		rsp.Msg = "email:" + err1.String()
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
	err1 = gvalid.Check(g.PostForm("phone"), "required|phone", "phone")
	if err1 != nil {
		rsp.Msg = err1.String()
		rsp.Code = 201
		rsp.Data = err1.Maps()
		g.JSON(http.StatusOK, rsp)
		return
	}

	mgo := db.InitMongoDB()
	newuser.Id = primitive.NewObjectID()
	gender, err := strconv.Atoi(g.PostForm("gender"))
	if err == nil {
		newuser.Gender = gender
	}
	newuser.Email = g.PostForm("email")
	newuser.Username = g.PostForm("username")
	newuser.Password = g.PostForm("password")
	newuser.Phone = g.PostForm("phone")
	newuser.Address = g.PostForm("address")

	//2. 采用CheckStruct 方式校验
	rules := map[string]string{
		"Phone": "required|phone",
		"Email": "email",
	}
	err1 = gvalid.CheckStruct(newuser, rules, nil)

	if err1 != nil {
		rsp.Msg = err1.String()
		rsp.Code = 201
		rsp.Data = err1.Maps()
		g.JSON(http.StatusOK, rsp)
		return
	}

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
	mgo := db.InitMongoDB()
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

/*  模糊查询 query  by username */
func QueryByUsername(g *gin.Context) {
	rsp := new(Rsp)
	fmt.Println(".....QueryByUsername..")
	username := g.Query("username")
	mgo := db.InitMongoDB()
	fmt.Println(username)

	filter := bson.M{"username": bson.M{"$regex": username, "$options": "$i"}}

	cur, err := mgo.Collection(db.User).Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
	}
	var users []User
	for cur.Next(context.Background()) {
		elme := new(User)
		err := cur.Decode(elme)
		if err == nil {
			users = append(users, *elme)
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
	log.Println("Getalluser.........")
	rsp := new(Rsp)
	mgo := db.InitMongoDB()
	filter := bson.M{}
	limit, err := strconv.Atoi(g.Query("limit"))
	page, err := strconv.Atoi(g.Query("page"))
	fmt.Println(page)

	//排序 正序1 倒序-1  ----------------------------
	opts := new(options.FindOptions)
	sortMap := make(map[string]interface{})
	sortMap["gender"] = -1
	opts.Sort = sortMap
	//排序 正序1 倒序-1  ----------------------------

	var users []User
	cur, err := mgo.Collection(db.User).Find(context.Background(), filter, opts.SetLimit(int64(limit)))
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
	mgo := db.InitMongoDB()

	var err1 *gvalid.Error
	//1. 单参数校验
	err1 = gvalid.Check(g.PostForm("gender"), "integer|between:0,2", nil)
	if err1 != nil {
		rsp.Msg = err1.String()
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}

	id := g.PostForm("id")

	oldId, err := primitive.ObjectIDFromHex(id)
	newpassword := g.PostForm("password")
	newgender, err := strconv.Atoi(g.PostForm("gender"))

	user := new(User)
	user.Id = oldId

	oldInfo := mgo.Collection(db.User).FindOne(context.Background(), bson.M{"_id": oldId})
	if oldInfo != nil {
		oldInfo.Decode(user)
	}
	if g.PostForm("username") != "" {
		user.Username = g.PostForm("username")
	}
	if newpassword != "" {
		user.Password = newpassword
	}
	if newgender != 0 {
		user.Gender = newgender
	}
	if g.PostForm("address") != "" {
		user.Address = g.PostForm("address")
	}

	//info, err := mgo.Collection(db.User).UpdateOne(context.Background(), bson.M{"_id": bsonx.ObjectID(oldId)},
	//	bson.M{"$set": bson.M{"username": newUsername}})
	//
	info := mgo.Collection(db.User).FindOneAndReplace(context.Background(), bson.M{"_id": oldId},
		user)
	if info.Err() != nil {
		rsp.Msg = info.Err().Error()
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
	fmt.Println(info)
	if err == nil {
		//fmt.Println(info.MatchedCount)
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
	mgo := db.InitMongoDB()
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

/* check user is exist ? */
func GetUserByNameAndPassword(username string, password string) bool {
	mgo := db.InitMongoDB()
	findfilter := bson.D{{"username", username}, {"password", password}}
	cur, err := mgo.Collection(db.User).Find(context.Background(), findfilter)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(User)
			err := cur.Decode(elme)
			if err == nil {
				return true
			}
		}
	}
	return false
}
