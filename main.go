package main

import (
	"bstgo-blog/config"
	"bstgo-blog/model"
	"bstgo-blog/router/home"
	"context"
	"log"
	"net/http"
	"time"

	"bstgo-blog/router/admin"

	mongocli "bstgo-blog/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func GroupRouterAdminMiddle(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	log.Println("=====================admin group router middle")
	//判断cookie中是否有session_id
	sessionId, err := c.Cookie(model.CookieSession)
	if err != nil {
		//没有sessionId则返回登录页面
		log.Println("no cookie sessionId ,return login")
		return
	}

	//连接数据库
	db := mongocli.MongoClient.Database(config.TotalCfgData.Mongo.Database)
	//指定连接集合
	col := db.Collection("session")
	//根据sessionId去数据库查找对应session
	session := model.Session{}
	err = col.FindOne(ctx, bson.M{"sid": sessionId}).Decode(&session)
	if err != nil {
		log.Println("session ", sessionId, "not found, return login")
		return
	}
	log.Println("session data is : ", session)
	c.Next()
}

func main() {
	mongocli.MongoInit()
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

	// 创建管理路由组
	adminGroup := router.Group("/admin")
	adminGroup.Use(GroupRouterAdminMiddle)
	{
		//管理首页
		adminGroup.GET("/", admin.Admin)
		//管理分类
		adminGroup.POST("/category", admin.Category)
		//管理排序页面，拖动标题达到排序效果
		adminGroup.POST("/sort", admin.Sort)
		//排序页面保存
		adminGroup.POST("/sortsave", admin.SortSave)
		//点击回到首页，管理页面首页
		adminGroup.POST("/index", admin.IndexList)
		// 创建分类
		adminGroup.POST("/createctg", admin.CreateCtg)
		// 创建子分类
		adminGroup.POST("/createsubctg", admin.CreateSubCtg)
		// 文章编辑界面
		adminGroup.GET("/articledit", admin.ArticleEdit)
		// 文章编辑发布
		adminGroup.POST("/pubarticle", admin.ArticlePub)
	}

	router.Run(":8080")
	mongocli.MongoRelease()
}
