package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"myGoProjectNew/db"
	"net/http"
)

type TestC struct {
	Mgo *mongo.Database
	//RedisCli *redis.Client
}

func (t TestC) Test3(g *gin.Context) {

	cur, err := t.Mgo.Collection(db.Test2).Indexes().List(context.Background())
	t.Mgo.Collection(db.Test2).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bsonx.Doc{{"height", bsonx.Int32(1)}, {"chainid", bsonx.Int32(1)}}, Options: options.Index().SetUnique(true),
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Print(cur)
}

//func (t TestC) TestInsert(g *gin.Context) {
//	rsp := new(Rsp)
//
//	var tests []interface{}
//	for i := 0; i < 1000000; i++ {
//		newuser := new(Test)
//
//		newuser.Email = strconv.Itoa(i)
//		newuser.Username = strconv.Itoa(i)
//		newuser.Password = strconv.Itoa(i)
//		newuser.Phone = strconv.Itoa(i)
//		newuser.Address = strconv.Itoa(i)
//		newuser.CreateTime = int32(time.Now().Unix())
//		newuser.Height = int32(i)
//		tests = append(tests, newuser)
//	}
//
//	insertID, err := t.Mgo.Collection(db.Test).InsertMany(context.Background(), tests, nil)
//	fmt.Println(insertID)
//	if err == nil {
//		rsp.Msg = "success"
//		rsp.Code = 200
//		g.JSON(http.StatusOK, rsp)
//		return
//	} else {
//		rsp.Msg = "faild"
//		rsp.Code = 201
//		g.JSON(http.StatusOK, rsp)
//		return
//	}
//}
func (t TestC) Test1Insert(g *gin.Context) {
	//rsp := new(Rsp)
	//
	//var tests []interface{}
	//
	//newuser := new(Test)
	//
	//newuser.Email = "sg"
	//newuser.Username = "hhhhh"
	//newuser.Password = "hhh"
	//newuser.Phone =  "hhhhh"
	//newuser.Address =  "hhhhh"
	//newuser.Height = 123456
	//tests = append(tests, newuser)
	//
	//insertID, err := t.Mgo.Collection(db.Test1).FindOneAndUpdate(context.Background(), tests, )
	//fmt.Println(insertID)
	//if err == nil {
	//	rsp.Msg = "success"
	//	rsp.Code = 200
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//} else {
	//	rsp.Msg = "faild"
	//	rsp.Code = 201
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//}
}
func (t TestC) Test2(g *gin.Context) {

	//var par = g.Query("abc")
	//fmt.Println(par)
	//
	//var Tests = make([]Test, 0)
	//rsp := new(Rsp)
	//
	//opts := new(options.FindOptions)
	//limit := int64(6)
	//skip := int64((1 - 1) * 6) //
	//opts.Limit = &limit
	//opts.Skip = &skip
	//
	////sortMap := make(map[string]interface{})
	////sortMap["height"] = -1
	////opts.Sort = sortMap
	//
	//filter := bson.D{{}}
	//Txs, err := t.Mgo.Collection(db.Test).Find(context.Background(), filter)
	//if err != nil {
	//	rsp.Code = 500
	//	rsp.Msg = "get data error"
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//}
	//for Txs.Next(context.Background()) {
	//	elem := new(Test)
	//	err := Txs.Decode(elem)
	//	if err != nil {
	//		//log.Error(err)
	//	}
	//	fmt.Println(elem.Id.Hex())
	//	Tests = append(Tests, *elem)
	//}
	//rsp.Data = Tests
	//rsp.Code = 200
	//rsp.Msg = "success"
	//g.JSON(http.StatusOK, rsp)
	//
	//return
}

func (t TestC) GetTestInfo(g *gin.Context) {
	rsp := new(Rsp)
	rsp.Msg = "faild"
	rsp.Code = 201
	g.JSON(http.StatusOK, rsp)
	return
}
func (t TestC) GetTestAddMyql(g *gin.Context) {
	resultList := new(ResultAccountList)
	db := db.InitMysqlDb()
	rows, err := db.Query("select * from account;")
	if err == nil {
		fmt.Println(rows)

		// for循环
		for rows.Next() {
			accountInfo := new(AccountInfo)
			errS := rows.Scan(&accountInfo.Id, &accountInfo.UserName, &accountInfo.Password)
			if errS != nil {
				fmt.Print(errS.Error())
			}
			resultList.DataList = append(resultList.DataList, *accountInfo)
		}
		fmt.Print(resultList)
	}

	rsp := new(Rsp)
	rsp.Msg = "faild"
	rsp.Code = 201
	rsp.Data = resultList
	g.JSON(http.StatusOK, rsp)
	return
}

type ResultAccountList struct {
	DataList []AccountInfo `json:"dataList"`
}
type AccountInfo struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
