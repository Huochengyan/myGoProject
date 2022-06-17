package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myGoProjectNew/db"
	"myGoProjectNew/models"
	"net/http"
)

/**
获取PC运行情况
*/
func (m UserC) GetPcResourceList(g *gin.Context) {
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
	cur, err := m.Mgo.Collection(db.PcResource).Find(context.Background(), filter, opts)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(models.PcResource)
			err := cur.Decode(elme)
			if err == nil {
				users = append(users, *elme)
			}
		}
	}

	sCount, _ := m.Mgo.Collection(db.PcResource).CountDocuments(context.Background(), filter)

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
新加 电脑使用情况
*/

func (m UserC) AddPcResource(g *gin.Context) {
	rsp := new(Rsp)
	var info models.PcResource
	err := g.BindJSON(&info)
	if err != nil {
		rsp.Msg = "json faild"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}

	//filter := bson.D{{"_id", info.Id}}
	//selector := bson.M{"_id": updateId}
	//updateInfo,_ :=bson.Marshal(&info)
	info.Id = primitive.NewObjectID()
	result, err := m.Mgo.Collection(db.PcResource).InsertOne(context.Background(), info)
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

func (m UserC) AddPcResourceInfo(info models.PcResource) {
	//filter := bson.D{{"_id", info.Id}}
	//selector := bson.M{"_id": updateId}
	//updateInfo,_ :=bson.Marshal(&info)
	info.Id = primitive.NewObjectID()
	result, err := m.Mgo.Collection(db.PcResource).InsertOne(context.Background(), info)
	if err == nil {
		fmt.Println(result.InsertedID)
	}

}
