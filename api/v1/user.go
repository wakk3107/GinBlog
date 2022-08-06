package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"GinBlog/utils/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var code int
	err := c.ShouldBindJSON(&data)
	msg, code := validator.Validate(data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	if err != nil {
		fmt.Println("用户绑定失败：", err)
		code = errmsg.ERROR_USER_TYPE
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询用户
func GetUser(c *gin.Context) {
	id, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errmsg.GetErrMsg(errmsg.ERROR_REQUEST),
		})
		return
	}
	data, code := model.GetUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询用户列表
func GetUsers(c *gin.Context) {
	PageSize, err1 := strconv.Atoi(c.Query("pagesize"))
	PageNum, err2 := strconv.Atoi(c.Query("pagenum"))
	UserName := c.Query("username")
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errmsg.GetErrMsg(errmsg.ERROR_REQUEST),
		})
		return
	}
	if PageSize == 0 {
		PageSize = -1
	}
	if PageNum == 0 {
		PageNum = -1
	}

	data, total := model.GetUsers(UserName, PageSize, PageNum)

	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.SUCCESS,
			"total":   total,
			"data":    data,
			"message": errmsg.GetErrMsg(errmsg.SUCCESS),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.SUCCESS,
			"total":   total,
			"data":    "null",
			"message": errmsg.GetErrMsg(errmsg.SUCCESS),
		})
	}

}

//编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("用户绑定失败：", err)
		code = errmsg.ERROR_USER_TYPE
	}
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(data)
	code = model.CheckUserEdit(uint(id), data.UserName)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
