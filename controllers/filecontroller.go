package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

/* upload file */
func Uploadfile(g *gin.Context) {
	fmt.Println("......uploadfile")
	file, header, err := g.Request.FormFile("file")
	if err == nil {
		fmt.Println("...", file)
		filename := header.Filename
		fmt.Println(file, err, filename)

		var uploadir string
		uploadir = "upload/11/"
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
		g.JSON(http.StatusOK, rsp)
		return
	} else {
		fmt.Println("..err..", err)
	}
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
