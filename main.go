package main

import (
	"bstgo-blog/router/home"
	"net/http"

	"bstgo-blog/router/admin"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(Cors()) //默认跨域
	//加载模板文件
	router.LoadHTMLGlob("views/**/*")
	//设置资源共享目录
	router.StaticFS("/static", http.Dir("./public"))
	//用户浏览首页
	router.GET("/home", home.Home)
	//用户浏览你分类
	router.GET("/category", home.Category)
	//管理首页
	router.GET("/admin", admin.Admin)
	//管理分类
	router.POST("/admin/category", admin.Category)
	//管理排序页面，拖动标题达到排序效果
	router.POST("/admin/sort", admin.Sort)
	//排序页面保存
	router.POST("/admin/sortsave", admin.SortSave)
	//点击回到首页，管理页面首页
	router.POST("/admin/index", admin.IndexList)
	// 创建分类
	router.POST("/admin/createctg", admin.CreateCtg)
	// 创建子分类
	router.POST("/admin/createsubctg", admin.CreateSubCtg)
	// 文章编辑
	router.POST("/admin/article-edit", admin.ArticleEdit)

	router.Run(":8080")
}
