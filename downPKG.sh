#查看配置
go env

#设置 使用七牛云下载
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go env

#引用项目需要的依赖增加到go.mod文件。
#去掉go.mod文件中项目不需要的依赖。
#开始下载
go mod tidy


