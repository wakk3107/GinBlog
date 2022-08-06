package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("用户绑定失败：", err)
		code = errmsg.ERROR_CATE_TYPE
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
			"err":     err,
		})
		return
	}

	code = model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//TODO 查询单个文章信息
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//TODO 查询文章列表
func GetArticle(c *gin.Context) {
	PageSize, err1 := strconv.Atoi(c.Query("pagesize"))
	PageNum, err2 := strconv.Atoi(c.Query("pagenum"))
	tittle := c.Query("tittle")
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

	data, code, total := model.GetArticle(tittle, PageSize, PageNum)

	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"total":   total,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"total":   total,
			"data":    "null",
			"message": errmsg.GetErrMsg(code),
		})
	}

}

//TODO 查询分类下所有文章
func GetCateArt(c *gin.Context) {
	PageSize, err1 := strconv.Atoi(c.Query("pagesize"))
	PageNum, err2 := strconv.Atoi(c.Query("pagenum"))
	cid, err3 := strconv.Atoi(c.Param("cid"))
	if err1 != nil || err2 != nil || err3 != nil {
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

	data, code, total := model.GetCateArt(cid, PageSize, PageNum)

	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"total":   total,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"total":   total,
			"data":    "null",
			"message": errmsg.GetErrMsg(code),
		})
	}
}

//编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("用户绑定失败：", err)
		code = errmsg.ERROR_CATE_TYPE
	}

	id, _ := strconv.Atoi(c.Param("id"))
	code = model.EditArticle(id, &data)

	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DelArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
func HandleGetAllData(c *gin.Context) {
	//log.Print("handle log")
	fmt.Println("hhhh")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("---body/--- \r\n " + string(body))

	fmt.Println("---header/--- \r\n")
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	//fmt.Println("header \r\n",c.Request.Header)

	c.JSON(200, gin.H{
		"receive": "1024",
	})

}
