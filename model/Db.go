package model

import (
	"GinBlog/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error //不知道为什么要个全局的，不要的话获取不到 db

func InitDb() {
	utils.Init()
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPass,
		utils.DbName)
	db, err = gorm.Open(utils.Db, dsn)
	if err != nil {
		fmt.Println("数据库连接失败，请检查参数\nError ", err)
		return
	}
	fmt.Println("数据库连接成功")
	db.SingularTable(true)

	db.AutoMigrate(&User{}, &Article{}, &Category{}, &Profile{}, &Comment{})
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 * time.Second)
}
