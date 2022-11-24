package controllers

import "myGoProjectNew/myProjectUtils"

//初始化

func InitServer() {
	println("初始化！！")
	config := myProjectUtils.GetConfInfo("qiqiu")
	QINIU_bucket = config.Key("bucket").String()
	QINIU_accessKey = config.Key("accessKey").String()
	QINIU_secretKey = config.Key("secretKey").String()
	QINIU_domainName = config.Key("domainName").String()
}

var (
	QINIU_bucket     string = ""
	QINIU_accessKey  string = ""
	QINIU_secretKey  string = ""
	QINIU_domainName        = ""
)
