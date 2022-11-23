package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"myGoProjectNew/db"
	"myGoProjectNew/models"
	"myGoProjectNew/pkg/util"
	"net/http"
	"strconv"
	"time"
)

type UserC struct {
	Mgo *mongo.Database
	//RedisCli *redis.Client
}

/*login user*/
func (m UserC) Login(g *gin.Context) {
	fmt.Println("login.........")
	rsp := new(Rsp)
	//name := g.PostForm("username")
	//pass := g.PostForm("password")

	var info models.User
	err := g.BindJSON(&info)
	if err != nil {
		rsp.Msg = "json faild"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
	name := info.Username
	pass := info.Password
	//var gerr *gvalid.Error
	//gerr = gvalid.Check(g.PostForm("username"), "required", nil)
	//if gerr != nil {
	//	rsp.Msg = "faild"
	//	rsp.Code = 201
	//	rsp.Data = gerr.Maps()
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//}
	//gerr = gvalid.Check(g.PostForm("password"), "required", nil)
	//if gerr != nil {
	//	rsp.Msg = "faild"
	//	rsp.Code = 201
	//	rsp.Data = gerr.Maps()
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//}

	findfilter := bson.D{{"username", name}, {"password", pass}}
	cur, err := m.Mgo.Collection(db.User).Find(context.Background(), findfilter)
	if err != nil {
		rsp.Msg = "faild"
		rsp.Code = 201
		rsp.Data = err.Error()
		g.JSON(http.StatusOK, rsp)
		return
	}
	for cur.Next(context.Background()) {
		elme := new(models.User)
		err := cur.Decode(elme)
		if err == nil {
			if elme.Username == name && elme.Password == pass {
				var info = new(LoginInfo)
				info.User = elme
				token, err := util.GenerateToken(g.PostForm("username"), g.PostForm("password"))
				if err == nil {
					info.Token = token
				}
				var loglogin models.LogLogin
				loglogin.Ip = g.ClientIP()
				loglogin.Username = elme.Username
				addLoginLog(m, loglogin)
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

/* query all user */
func (m UserC) Queryalluser(g *gin.Context) {
	fmt.Println("Queryalluser.........")
	rsp := new(Rsp)
	var users []models.User
	cur, err := m.Mgo.Collection(db.User).Find(context.Background(), bson.D{}, nil)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(models.User)
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
func (m UserC) Getalluser(g *gin.Context) {
	log.Println("Getalluser.........")
	rsp := new(Rsp)

	filter := bson.M{}

	//limit, err := strconv.Atoi(g.Query("limit"))
	page, err := strconv.Atoi(g.Query("page"))
	fmt.Println(page)

	//排序 正序1 倒序-1  ----------------------------
	opts := new(options.FindOptions)
	sortMap := make(map[string]interface{})
	sortMap["gender"] = -1
	opts.Sort = sortMap
	//opts.Limit=int64(limit)
	//排序 正序1 倒序-1  ----------------------------

	var users []models.User
	cur, err := m.Mgo.Collection(db.User).Find(context.Background(), filter, opts)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(models.User)
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

/**
获取用户列表
*/
func (m UserC) GetUserList(g *gin.Context) {
	rsp := new(Rsp)
	//pageIndex := g.PostForm("pageIndex")
	//pageSize := g.PostForm("pageSize")
	//fmt.Println(pageIndex+pageSize)

	var info models.PageHelper
	err := g.BindJSON(&info)
	if err != nil {
		rsp.Msg = "json faild"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
	page := info.PageIndex
	pageSize := info.PageSize

	/* 默认 */
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	opts := new(options.FindOptions)
	limit := int64(pageSize)
	skip := int64((page - 1) * pageSize)
	opts.Limit = &limit
	opts.Skip = &skip

	filter := bson.M{}

	//limit, err := strconv.Atoi(g.Query("limit"))
	//page, err := strconv.Atoi(g.Query("page"))
	//fmt.Println(page)

	//排序 正序1 倒序-1  ----------------------------
	//opts := new(options.FindOptions)
	//sortMap := make(map[string]interface{})
	//sortMap["gender"] = -1
	//opts.Sort = sortMap
	//opts.Limit=int64(limit)
	//排序 正序1 倒序-1  ----------------------------

	//var users []models.User
	users := make([]interface{}, 0)
	cur, err := m.Mgo.Collection(db.User).Find(context.Background(), filter, opts)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(models.User)
			err := cur.Decode(elme)
			if err == nil {
				users = append(users, *elme)
			}
		}
	}

	sCount, _ := m.Mgo.Collection(db.User).CountDocuments(context.Background(), filter)

	var resultData models.UserList
	resultData.Data = users
	resultData.Total = sCount

	rsp.Msg = "success"
	rsp.Code = 0
	rsp.Data = resultData
	g.JSON(http.StatusOK, rsp)
	return
}

/**
修改用户列表
*/
func (m UserC) Update(g *gin.Context) {
	rsp := new(Rsp)
	var info models.User
	err := g.BindJSON(&info)
	if err != nil {
		rsp.Msg = "json faild"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}

	if info.Id.String() == "" {
		rsp.Msg = "id is empty!"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}

	filter := bson.D{{"_id", info.Id}}
	//selector := bson.M{"_id": updateId}
	//updateInfo,_ :=bson.Marshal(&info)
	result := m.Mgo.Collection(db.User).FindOneAndReplace(context.Background(), filter, info)
	if result != nil {
		if result.Err() != nil {
			fmt.Println(result.Err())
		}
		rsp.Msg = "success"
		rsp.Code = 200
		rsp.Data = 1
		g.JSON(http.StatusOK, rsp)
		return
	}
	rsp.Msg = "err"
	rsp.Code = 201
	rsp.Data = 0
	g.JSON(http.StatusOK, rsp)
	return

}

/**
删除用户信息
*/

func (m UserC) DelUser(g *gin.Context) {
	rsp := new(Rsp)
	var info models.User
	err := g.BindJSON(&info)
	if err != nil {
		rsp.Msg = "json faild"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}

	if info.Id.String() == "" {
		rsp.Msg = "id is empty!"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}

	filter := bson.D{{"_id", info.Id}}
	//selector := bson.M{"_id": updateId}
	//updateInfo,_ :=bson.Marshal(&info)
	result := m.Mgo.Collection(db.User).FindOneAndDelete(context.Background(), filter)
	if result != nil {
		if result.Err() != nil {
			fmt.Println(result.Err())
		}
		rsp.Msg = "success"
		rsp.Code = 200
		rsp.Data = 1
		g.JSON(http.StatusOK, rsp)
		return
	}
	rsp.Msg = "err"
	rsp.Code = 201
	rsp.Data = 0
	g.JSON(http.StatusOK, rsp)
	return

}

/**
新加 用户信息
*/

func (m UserC) AddUser(g *gin.Context) {
	rsp := new(Rsp)
	var info models.User
	err := g.BindJSON(&info)
	if err != nil {
		rsp.Msg = "json faild"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}

	if info.Id.String() == "" {
		rsp.Msg = "id is empty!"
		rsp.Code = 201
		rsp.Data = 0
		g.JSON(http.StatusOK, rsp)
		return
	}

	//filter := bson.D{{"_id", info.Id}}
	//selector := bson.M{"_id": updateId}
	//updateInfo,_ :=bson.Marshal(&info)
	info.Id = primitive.NewObjectID()
	result, err := m.Mgo.Collection(db.User).InsertOne(context.Background(), info)
	if err == nil {
		fmt.Println(result.InsertedID)
		rsp.Msg = "success"
		rsp.Code = 200
		rsp.Data = 1
		g.JSON(http.StatusOK, rsp)
		return
	}
	rsp.Msg = "err"
	rsp.Code = 201
	rsp.Data = 0
	g.JSON(http.StatusOK, rsp)
	return

}

func addLoginLog(m UserC, info models.LogLogin) {
	info.Id = primitive.NewObjectID()
	info.CreateTime = time.Now().Unix()
	result, err := m.Mgo.Collection(db.LogLogin).InsertOne(context.Background(), info)
	if err == nil {
		fmt.Println(result.InsertedID)
	}

	//如果获取的客户端IP为 127.0.0.1 配置下Nginx
	//location / {
	//	proxy_pass http://127.0.0.1:8080;
	//	proxy_set_header X-Real-IP $remote_addr;
	//	proxy_set_header X-Forward-For $remote_addr;
	//}
}
