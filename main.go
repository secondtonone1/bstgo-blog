package main

import (
	"bstgo-blog/model"
	"bstgo-blog/router/home"
	"log"
	"net/http"

	"bstgo-blog/router/admin"

	mongocli "bstgo-blog/mongo"

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

func GroupRouterAdminMiddle(c *gin.Context) {

	log.Println("=====================admin group router middle")
	//判断cookie中是否有session_id
	sessionId, err := c.Cookie(model.CookieSession)
	if err != nil {
		//没有sessionId则返回登录页面
		log.Println("no cookie sessionId ,return login")
		c.HTML(http.StatusOK, "admin/login.html", nil)
		c.Abort()
		return
	}
	sessionData, err := mongocli.GetSessionById(sessionId)
	if err != nil {
		log.Println("get sessionid ", sessionId, "failed, return login")
		c.HTML(http.StatusOK, "admin/login.html", nil)
		c.Abort()
		return
	}
	log.Println("session data is : ", sessionData)
	c.Next()
}

func CheckLogin(c *gin.Context) {
	log.Println("check login midware")
	//判断cookie中是否有session_id
	sessionId, err := c.Cookie(model.CookieSession)
	if err != nil {
		//没有sessionId则返回登录页面
		log.Println("no cookie sessionId ,return login")
		baseRsp := model.BaseRsp{}
		baseRsp.Code = model.ERR_NO_LOGIN
		baseRsp.Msg = model.MSG_NO_LOGIN
		c.JSON(http.StatusOK, baseRsp)
		c.Abort()
		return
	}
	sessionData, err := mongocli.GetSessionById(sessionId)
	if err != nil {
		log.Println("get sessionid ", sessionId, "failed, return login")
		baseRsp := model.BaseRsp{}
		baseRsp.Code = model.ERR_NO_LOGIN
		baseRsp.Msg = model.MSG_NO_LOGIN
		c.JSON(http.StatusOK, baseRsp)
		c.Abort()
		return
	}
	log.Println("session data is : ", sessionData)
	c.Next()
}

func CalCulateVisit() gin.HandlerFunc {
	return func(c *gin.Context) {
		num, err := mongocli.AddVisitNum()
		if err != nil {
			log.Println("add visit num failed, err is ", err)
			return
		}

		c.Set("visitnum", num)
	}
}

func GetVisitMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		num, err := mongocli.GetVisitNum()
		if err != nil {
			log.Println("add visit num failed, err is ", err)
			return
		}

		c.Set("visitnum", num)
	}
}

func main() {
	mongocli.MongoInit()
	router := gin.Default()
	router.Use(Cors()) //默认跨域
	//加载模板文件
	router.LoadHTMLGlob("views/**/*")
	//设置资源共享目录
	router.StaticFS("/static", http.Dir("./public"))

	pageGroup := router.Group("/")
	pageGroup.Use(CalCulateVisit())
	{
		//用户浏览首页
		pageGroup.GET("/home", CalCulateVisit(), home.Home)
		//用户浏览你分类
		pageGroup.GET("/category", CalCulateVisit(), home.Category)

		//用户浏览单个文章
		pageGroup.GET("/articlepage", CalCulateVisit(), home.ArticlePage)
	}

	homeGroup := router.Group("/home")

	//新增用户点赞数
	homeGroup.POST("/addlovenum", home.AddLoveNum)

	//点击用户评论
	homeGroup.POST("/comment", home.Comment)

	//点击评论喜欢
	homeGroup.POST("/addcomlove", home.ComLove)
	//评论回复
	homeGroup.POST("/comreply", home.ComReply)
	//回复区点赞
	homeGroup.POST("/addrpllove", home.ReplyLove)

	//请求子分类下文章信息列表
	homeGroup.POST("/artinfos", home.SubCatArtInfos)

	//admin登录页面
	router.GET("/admin/login", admin.Login)
	//admin 登录提交
	router.POST("/admin/loginsub", admin.LoginSub)

	// 创建管理路由组
	adminGroup := router.Group("/admin")
	adminGroup.Use(GroupRouterAdminMiddle, GetVisitMid())
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
		// 子标签排序
		adminGroup.POST("/sortmenu", admin.SortMenu)
		// 文章编辑界面
		adminGroup.GET("/articledit", admin.ArticleEdit)
		//获取子分类下拉菜单
		adminGroup.POST("/subcatselect", admin.SubCatSelect)
		//文章搜索返回列表
		adminGroup.POST("/articlesearch", admin.ArticleSearch)
		//文章删除
		adminGroup.POST("/delarticle", admin.DelArticle)
		//文章编辑
		adminGroup.GET("/articlemodify", admin.ModifyArticle)
		//文章更新
		adminGroup.POST("/updatearticle", admin.UpdateArticle)
	}

	// 文章编辑发布
	router.POST("admin/pubarticle", CheckLogin, admin.ArticlePub)
	//	mongocli.SetVisitNum()
	router.Run(":8080")
	mongocli.MongoRelease()
}
