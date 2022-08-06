package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode   string
	HttpPort  string
	JwtKey    string
	Db        string
	DbHost    string
	DbPort    string
	DbUser    string
	DbPass    string
	DbName    string
	AccessKey string
	SecretKey string
	Bucket    string
	Url       string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	LoadServer(file)
	LoadDataBase(file)
	LoadQiniu(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("awrawea")
}
func LoadDataBase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbName = file.Section("database").Key("DbName").MustString("ginlog")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPass = file.Section("database").Key("DbPass").MustString("245330")
}
func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("t-bRhW2JdUwaETsZzeFkaVIaqzATr3gFEFUyxmnN")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("jByGU5XBaUaPBO2BQWYHOwjBaKkS8k0zLqWK67qo")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("wakk3107")
	Url = file.Section("qiniu").Key("QiniuServer").MustString("http://rfd843vfl.hd-bkt.clouddn.com/")
}
