package model

import (
	"GinBlog/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	ID        uint      `gorm:"primary_key"`
	Category  Category  `gorm:"foreignkey:cid"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
	Tittle    string    `gorm:"type:varchar(100);not null" json:"tittle"`
	Cid       int       `gorm:"type:int;not null" json:"cid"`
	Desc      string    `gorm:"type:varchar(200)" json:"desc"`
	Content   string    `gorm:"type:longtext" json:"content"`
	Img       string    `gorm:"type:varchar(100)" json:"img"`
}

//查询文章是否存在
func CheckArticle(name string) int {
	var art Article
	db.Select("id").Where("name=?", name).First(&art)
	return errmsg.SUCCESS
}

//新增文章
func CreateArticle(Data *Article) int {
	err := db.Create(&Data).Error
	if err != nil {
		fmt.Println("用户添加失败：", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询文章列表
func GetArticle(tittle string, PageSize int, PageNum int) ([]Article, int, int) {
	var artList []Article
	var total = 0
	offset := (PageNum - 1) * PageSize
	if PageNum == -1 && PageSize == -1 {
		offset = -1
	}
	if tittle != "" {
		err := db.Preload("Category").Where("tittle=?", tittle).Limit(PageSize).Offset(offset).Find(&artList).Count(&total).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		}
	} else {
		err := db.Preload("Category").Limit(PageSize).Offset(offset).Find(&artList).Count(&total).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		}
	}
	return artList, errmsg.SUCCESS, total

}

func GetCateArt(cid int, PageSize int, PageNum int) ([]Article, int, int) {
	var artList []Article
	var total = 0
	offset := (PageNum - 1) * PageSize
	if PageNum == -1 && PageSize == -1 {
		offset = -1
	}
	err := db.Preload("Category").Limit(PageSize).Offset(offset).Where("cid=?", cid).Find(&artList).Count(&total).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCESS, total

}

//删除文章
func DelArticle(Id int) int {
	var art Article
	err := db.Where("Id=?", Id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑文章
func EditArticle(Id int, Data *Article) int {
	var maps = make(map[string]interface{})
	maps["tittle"] = Data.Tittle
	maps["cid"] = Data.Cid
	maps["category"] = Data.Category
	maps["desc"] = Data.Desc
	maps["content"] = Data.Content
	maps["img"] = Data.Img
	err := db.Model(&Article{}).Where("id=?", Id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询单个文章信息
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_FOUND
	}
	return art, errmsg.SUCCESS
}
