package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("用户绑定失败：", err)
		code = errmsg.ERROR_CATE_TYPE
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类
func GetOneCate(c *gin.Context) {
	cid, err1 := strconv.Atoi(c.Param("cid"))
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errmsg.GetErrMsg(errmsg.ERROR_REQUEST),
		})
		return
	}
	data, code := model.GetOneCate(cid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

//查询分类列表
func GetCategorys(c *gin.Context) {
	PageSize, err1 := strconv.Atoi(c.Query("pagesize"))
	PageNum, err2 := strconv.Atoi(c.Query("pagenum"))
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

	data, total := model.GetCategory(PageSize, PageNum)

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

//编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("用户绑定失败：", err)
		code = errmsg.ERROR_CATE_TYPE
	}
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		id, _ := strconv.Atoi(c.Param("id"))
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DelCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

//TODO 查询分类下所有文章
