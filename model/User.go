package model

import (
	"GinBlog/utils/errmsg"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
	UserName  string    `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	PassWord  string    `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role      int       `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=1" label:"角色码"`
}

//查询用户是否存在
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("user_name=?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

//查询修改状态下的用户是否存在
func CheckUserEdit(id uint, name string) int {
	var user User
	db.Select("*").Where("user_name=?", name).First(&user)
	if user.ID == id {
		return errmsg.SUCCESS //1001
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.ERROR
}

//新增用户
func CreateUser(Data *User) int {
	err := db.Create(&Data).Error
	if err != nil {
		fmt.Println("用户添加失败：", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Where("id=?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

//查询用户列表
func GetUsers(UserName string, PageSize int, PageNum int) ([]User, int) {
	var users []User
	var total = 0
	offset := (PageNum - 1) * PageSize
	if PageNum == -1 && PageSize == -1 {
		offset = -1
	}
	if UserName == "" {
		err := db.Offset(offset).Find(&users).Limit(PageSize).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
	}
	err := db.Where("user_name like ?", UserName+"%").Offset(offset).Find(&users).Limit(PageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	db.Model(&users).Count(&total)

	return users, total

}

//密码加密
func ScryptPW(PassWord string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 12, 32, 44, 55, 123, 22}
	HashPW, err := scrypt.Key([]byte(PassWord), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	fpw := base64.StdEncoding.EncodeToString(HashPW)
	return fpw

}

//钩子函数 存到数据库前先对密码进行加密
func (u *User) BeforeSave() {
	u.PassWord = ScryptPW(u.PassWord)
}

//删除用户
func DelUser(Id int) int {
	var user User
	err := db.Where("Id=?", Id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑用户
func EditUser(Id int, Data *User) int {
	var maps = make(map[string]interface{})
	maps["user_name"] = Data.UserName
	maps["role"] = Data.Role
	err := db.Model(&User{}).Where("id=?", Id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//登录验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("user_name=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPW(password) != user.PassWord {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.SUCCESS
}

//登录验证
func CheckFrontLogin(username string, password string) int {
	var user User
	db.Where("user_name=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPW(password) != user.PassWord {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}
