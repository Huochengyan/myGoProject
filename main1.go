package main11

import (
	"container/list"
	"context"
	_ "crypto/md5"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	//"github.com/mongodb/mongo-go-driver/bson/bsontype"
	"github.com/mongodb/mongo-go-driver/mongo"
	_ "github.com/mongodb/mongo-go-driver/mongo"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	_ "strconv"
	"strings"
	"time"
)

const tname  ="test"
/* mongodb */
func InitMongoDB() (collection *mongo.Database, err error){
	const url  = "mongodb://192.168.1.108:27017";
	const dbName="mycol"
	client,err :=mongo.NewClient(url)
	if err != nil{
		return nil,err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	collection = client.Database(dbName)
	return collection, nil
}

/* test hello */
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

/* index login page */
func index(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "login.html", http.StatusFound)
}
/* uploadfile */
func uploadfile(w http.ResponseWriter, r *http.Request)  {

	/* 文件。。 */
	file, _, err := r.FormFile("file")

	if err != nil {
		log.Fatal("FormFile: ", err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Close: ", err.Error())
			return
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("ReadAll: ", err.Error())
		return
	}

	w.Write(bytes)

}

type Testinfo struct {
	Title       string `json:"title"`
	CreateTime  int64 `json: "createtime"`
	Description int32 `json:"description"`
	likes        int32 `json:"likes"`
    By   	 	 int32 `json:"by"`
}
func main() {
	mgdb, errdb:=InitMongoDB()
	if errdb != nil {
		log.Fatal("init db error...........: ", errdb)
	}
	chainidFilter := bson.M{"by": 5}

	//静态文件处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/"))))
	//html文件处理  http://localhost:9090   index.html
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("views/"))))

	// [test] table name
    cur,err :=mgdb.Collection(tname).Find(context.Background(),chainidFilter)
    info :=new(Testinfo)
	err1:=mgdb.Collection(tname).FindOne(context.Background(),chainidFilter).Decode(info)
	if err1==nil{
		println(info.CreateTime)
	}
	var list []Testinfo
    if err==nil{
		for cur.Next(context.Background()){
			elem  :=new(Testinfo);
			err :=cur.Decode(elem)
			if err==nil {
				list = append(list,  *elem)
			}
		}
	}

	http.HandleFunc("/index/", index) //设置访问的路由
	http.HandleFunc("/hello/", sayhelloName) //设置访问的路由
	http.HandleFunc("/upload/", uploadfile) //设置访问的路由
	http.HandleFunc("/login/", login)


	lerrs := http.ListenAndServe(":9090", nil) //设置监听的端口
	if lerrs != nil {
		log.Fatal("ListenAndServe...........: ", lerrs)
	}


}



type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func login(w http.ResponseWriter, r *http.Request) {
 	fmt.Println("login.........")
}

func main1()  {
	fmt.Println("..............start")
	//
    //var numbers [] string
	//println(len(numbers))
	//
    ////append
	//numbers=append(numbers, "0")
	//numbers=append(numbers,"1")
	//numbers=append(numbers,"2","3","4","5")
	//println(len(numbers))
	//
	//index :=2
	//numbers=append(numbers[:index],numbers[index+1:]...)
	//println(len(numbers))

	var l list.List
    l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)

	println(l.Len())

	for i :=l.Front(); i!=nil;i=i.Next(){
		fmt.Println("value",i.Value)
	}

}
