package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	_ "github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	_ "github.com/qiniu/go-sdk/v7/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

type FileC struct {
	Mgo *mongo.Database
	//RedisCli *redis.Client
}

/* upload file */
func (m FileC) Uploadfile(g *gin.Context) {
	fmt.Println("......uploadfile")
	file, header, err := g.Request.FormFile("file")
	if err == nil {
		fmt.Println("...", file)
		filename := header.Filename
		fmt.Println(file, err, filename)

		var uploadir string
		uploadir = "./upload/"
		_, err := os.Stat(uploadir)
		if os.IsNotExist(err) {
			os.Mkdir(uploadir, os.ModePerm)
		}
		//创建文件
		out, err := os.Create(uploadir + filename)

		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		rsp := new(Rsp)
		rsp.Msg = "success"
		rsp.Code = 200
		rsp.Data = ""
		g.JSON(http.StatusOK, rsp)
		return
	} else {
		fmt.Println("..err..", err)
	}
}

func (m FileC) UploadFileQiNiu(g *gin.Context) {
	var accessKey = QINIU_accessKey
	var secretKey = QINIU_secretKey
	var uploadFilePath string
	var upLoadFileName string
	file, header, err := g.Request.FormFile("file")
	if err == nil {
		fmt.Println("...", file)
		filename := header.Filename
		var uploadir string
		uploadir = "./upload/"
		_, err := os.Stat(uploadir)
		if os.IsNotExist(err) {
			os.Mkdir(uploadir, os.ModePerm)
		}
		//创建文件
		out, err := os.Create(uploadir + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		uploadFilePath = uploadir + "/" + filename
		upLoadFileName = filename
	}
	localFile := uploadFilePath                 // "./upload/1aadb80d41ef98ec83acbe0a96f8591e.jpeg"
	bucket := QINIU_bucket                      // "hcyqiniu"
	key := "mygoproject" + "/" + upLoadFileName // "mygoproject/test.jpg"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		fmt.Println(err)
		return
	}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	rsp := new(Rsp)
	rsp.Msg = "success"
	rsp.Code = 200
	rsp.Data = ret
	g.JSON(http.StatusOK, rsp)
	return
}

/* download file  */
func Downloadfile(g *gin.Context) {
	g.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(path.Base("upload/1.jpg")))
	g.Header("Content-Description", "File Transfer")
	g.Header("Content-Type", "application/octet-stream")
	g.Header("Content-Transfer-Encoding", "binary")
	g.Header("Expires", "0")
	// 如果缓存过期了，会再次和原来的服务器确定是否为最新数据，而不是和中间的proxy
	g.Header("Cache-Control", "must-revalidate")
	g.Header("Pragma", "public")
	g.Header("content-disposition", `attachment; filename=`+"upload/1.jpg")
	g.File("upload/1.jpg")
}

/* read file  */
func Readfile(g *gin.Context) {
	b, err := ioutil.ReadFile("gopath.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	fmt.Println(str)
}
