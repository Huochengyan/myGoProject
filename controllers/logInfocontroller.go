package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myGoProjectNew/db"
	"myGoProjectNew/models"
	"net/http"
)

/**
获取用户登录日志
*/
func (m UserC) GetLoginLogList(g *gin.Context) {
	rsp := new(Rsp)
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

	//倒叙
	sortMap := make(map[string]interface{})
	sortMap["createtime"] = -1
	opts.Sort = sortMap

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
	cur, err := m.Mgo.Collection(db.LogLogin).Find(context.Background(), filter, opts)
	if err == nil {
		for cur.Next(context.Background()) {
			elme := new(models.User)
			err := cur.Decode(elme)
			if err == nil {
				users = append(users, *elme)
			}
		}
	}

	sCount, _ := m.Mgo.Collection(db.LogLogin).CountDocuments(context.Background(), filter)

	var resultData models.UserList
	resultData.Data = users
	resultData.Total = sCount

	rsp.Msg = "success"
	rsp.Code = 0
	rsp.Data = resultData
	g.JSON(http.StatusOK, rsp)
	return
}
