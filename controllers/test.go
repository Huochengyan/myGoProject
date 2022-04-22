package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myGoProjectNew/log"
	"strings"
	"time"
)

func (t TestC) Test4(g *gin.Context) {
	//rsp := new(Rsp)
	//newInfo:=new(ResultInfo)
	//newInfo.Id=primitive.NewObjectID()
	//
	//for i:=1;i<100;i++{
	//	time.Sleep(time.Second*2)
	//	fmt.Println("page: ",i)
	//    url:=""
	//	body:=utils.HttpGet(url)
	//	var result ResultDataOnline
	//	err:=json.Unmarshal(body,&result)
	//	if err!=nil{
	//		log.Error(err.Error())
	//		break
	//	}
	//	for j:=0;j<len(result.Data.Results) ;j++  {
	//		re, err := t.Mgo.Collection("test4").InsertOne(context.Background(), result.Data.Results[j])
	//		if err!=nil{
	//			log.Error(err.Error())
	//		}else{
	//			fmt.Println(re.InsertedID)
	//		}
	//	}
	//}
	//rsp.Data = newInfo
	//rsp.Code = 200
	//rsp.Msg = "success"
	//g.JSON(http.StatusOK, rsp)
	//return

	//
	//filter:=bson.D{{"status",2}}
	//Txs, err := t.Mgo.Collection("test3").Find(context.Background(), filter)
	//if err != nil {
	// log.Error(err.Error())
	//}
	//for Txs.Next(context.Background()) {
	//	elem := new(testRe)
	//	err := Txs.Decode(elem)
	//	if err != nil {
	//		log.Error(err.Error())
	//	}
	//	if elem.Transferresult==""{
	//		continue
	//	}
	//	info:=new(ResultDataChild)
	//	err=json.Unmarshal([]byte(elem.Transferresult),&info)
	//	if err!=nil{
	//		log.Error(err.Error())
	//	}
	//	filter := bson.D{{"_id", elem.Id}}
	//	update := bsonx.Doc{{"$set", bsonx.Document(bsonx.Doc{{"bill_id", bsonx.String(info.Bill_id)}})}}
	//	SingleResult := t.Mgo.Collection("test3").FindOneAndUpdate(context.Background(), filter, update)
	//	if SingleResult != nil {
	//		if SingleResult.Err() != nil {
	//			log.Error(SingleResult.Err().Error())
	//		}
	//	}
	//
	//	//reData:=new(ResultDataChild)
	//	//err=t.Mgo.Collection("test4").FindOne(context.Background(), filterId).Decode(&reData)
	//	//if err!=nil{
	//	//	if strings.Contains(err.Error(),"no documents in result"){
	//	//		fmt.Println(elem.Phone+"\t"+info.Bill_id)
	//	//		fmt.Println(info.Bill_id)
	//	//	}else{
	//	//		log.Error(err.Error())
	//	//	}
	//	//}
	//}

	filter := bson.D{{}}
	Txs, err := t.Mgo.Collection("test4").Find(context.Background(), filter)
	if err != nil {
		log.Error(err.Error())
	}
	for Txs.Next(context.Background()) {
		elem := new(ResultDataChild)
		err := Txs.Decode(elem)
		if err != nil {
			log.Error(err.Error())
		}
		filterId := bson.D{{"bill_id", elem.Bill_id}}
		reData := new(ResultDataChild)
		err = t.Mgo.Collection("test3").FindOne(context.Background(), filterId).Decode(&reData)
		if err != nil {
			if strings.Contains(err.Error(), "no documents in result") {
				fmt.Println(elem.Phone + "\t" + elem.Bill_id + "\t" + elem.Amount)
			} else {
				log.Error(err.Error())
			}
		}
	}

}

func (t TestC) Test5(g *gin.Context) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano() / 1e9)

	fmt.Println(time.Now().Format("20060102150405"))

	fmt.Println(time.Now().Format("20060102150405"))
}
func testdata(t *TestC) {
	filter := bson.D{{}}
	Txs, err := t.Mgo.Collection("test4").Find(context.Background(), filter)
	if err != nil {
		log.Error(err.Error())
	}
	for Txs.Next(context.Background()) {
		elem := new(ResultDataChild)
		err := Txs.Decode(elem)
		if err != nil {
			log.Error(err.Error())
		}
		filterId := bson.D{{"bill_id", elem.Bill_id}}
		reData := new(ResultDataChild)
		err = t.Mgo.Collection("test3").FindOne(context.Background(), filterId).Decode(&reData)
		if err != nil {
			if strings.Contains(err.Error(), "no documents in result") {
				fmt.Println(elem.Bill_id + "\t" + elem.Amount)
			} else {
				log.Error(err.Error())
			}
		}
	}
}

type ResultInfo struct {
	Id       primitive.ObjectID `bson:"_id"` //id
	Username string             `json:"username"`
	Password string             `json:"password"`
}

type ResultDataOnline struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    ResultDataTotal `json:"data"`
}
type ResultDataTotal struct {
	Count   int               `json:"code"`
	Results []ResultDataChild `json:"results"`
}
type ResultDataChild struct {
	Bill_id   string `json:"bill_id"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Amount    string `json:"amount"`
	Create_at string `json:"create_at"`
}

type testRe struct {
	Id             primitive.ObjectID `bson:"_id"` //id
	Phone          string             `json:"phone"`
	Transferresult string             `json:"transferresult"`
}
