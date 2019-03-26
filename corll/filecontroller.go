package corll

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
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
