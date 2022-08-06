package v1

import (
	"GinBlog/middleware"
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code := model.CheckLogin(data.UserName, data.PassWord)
	if code == errmsg.SUCCESS {
		token, code := middleware.GenToken(data.UserName)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"token":   token})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"token":   nil,
		})
	}

}
func LoginFront(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code := model.CheckLogin(data.UserName, data.PassWord)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"token":   nil,
		})
	}

}
