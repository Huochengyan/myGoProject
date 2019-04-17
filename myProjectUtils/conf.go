package myProjectUtils

import "github.com/go-ini/ini"

type MyPorjectConf struct {
}

/* 获取配置conf文件的参数信息 */
func GetConf(name string, key string) string {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		panic(err)
	}
	return cfg.Section(name).Key(key).String()
}
