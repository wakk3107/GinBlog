package routers

import (
	v1 "GinBlog/api/v1"
	"GinBlog/middleware"
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	//r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")
	r.Static("/front", "./web/front/dist")
	r.Static("/admin", "./web/admin/dist")
	r.LoadHTMLGlob("./web/front/dist/index.html")
	//不知道为啥，反正front的页面就是要这样才能加载进去
	r.GET("/front", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		//user模块的路由接口
		Auth.PUT("user/:id", v1.EditUser)
		Auth.DELETE("user/:id", v1.DelUser)
		//分类模块的路由接口
		Auth.POST("category/add", v1.AddCategory)
		Auth.PUT("category/:id", v1.EditCategory)
		Auth.DELETE("category/:id", v1.DelCategory)
		//文章模块的路由接口
		Auth.POST("article/add", v1.AddArticle)
		Auth.PUT("article/:id", v1.EditArticle)
		Auth.DELETE("article/:id", v1.DelArticle)
		//上传文件
		Auth.POST("upload", v1.Upload)
		// 评论模块
		Auth.GET("comment/list", v1.GetCommentList)
		Auth.DELETE("delcomment/:id", v1.DeleteComment)
		Auth.PUT("checkcomment/:id", v1.CheckComment)
		Auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}
	router := r.Group("api/v1")
	{
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		router.POST("user/add", v1.AddUser)
		router.GET("user/getusers", v1.GetUsers)
		router.GET("user/:id", v1.GetUser)

		router.GET("category/getcategorys", v1.GetCategorys)
		router.GET("category/:cid", v1.GetOneCate)

		router.GET("article/getarticles", v1.GetArticle)
		router.GET("article/:id", v1.GetArtInfo)
		router.GET("article/artbycate/:cid", v1.GetCateArt)
		router.POST("user/test", v1.HandleGetAllData)

		router.GET("profile/:id", v1.GetProfile)
		router.PUT("profile/:id", v1.UpdateProfile)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)
	}

	return r
}
