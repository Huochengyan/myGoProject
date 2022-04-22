package myProjectUtils

import (
	"fmt"
	"github.com/go-ini/ini"
)

/* 获取配置conf文件的参数信息 */
func GetConf(name string, key string) string {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println(err.Error())
	}
	return cfg.Section(name).Key(key).String()
}

var Config *ini.File

///* 获取配置conf文件的参数信息 */
func GetConfInfo(name string) *ini.Section {
	//cfg, err := ini.Load("conf/app.ini")
	cfg := Config
	return cfg.Section(name)
}
