package model

import (
	"GinBlog/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Category struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
	Name      string    `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) int {
	var Cate Category
	db.Select("id").Where("name=?", name).First(&Cate)
	if Cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCESS
}

//新增分类
func CreateCategory(Data *Category) int {
	err := db.Create(&Data).Error
	if err != nil {
		fmt.Println("用户添加失败：", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询单个分类
func GetOneCate(cid int) (Category, int) {
	var Cate Category
	err := db.Where("id=?", cid).Find(&Cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Cate, errmsg.ERROR
	}
	return Cate, errmsg.SUCCESS

}

//查询分类列表
func GetCategory(PageSize int, PageNum int) ([]Category, int) {
	var Cate []Category
	var total = 0
	offset := (PageNum - 1) * PageSize
	if PageNum == -1 && PageSize == -1 {
		offset = -1
	}
	err := db.Limit(PageSize).Offset(offset).Find(&Cate).Count(&total).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return Cate, total

}

//删除分类
func DelCategory(Id int) int {
	var Cate Category
	err := db.Where("Id=?", Id).Delete(&Cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑分类
func EditCategory(Id int, Data *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = Data.Name
	err := db.Model(&Category{}).Where("id=?", Id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
