package admin

import (
	"bstgo-blog/config"
	"bstgo-blog/model"
	mongocli "bstgo-blog/mongo"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"crypto/sha1"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func Admin(c *gin.Context) {

	var res string = model.RENDER_MSG_SUCCESS
	adminR := model.AdminIndexR{}
	adminR.Cur = 1
	adminR.Total = 1
	defer func() {
		c.HTML(http.StatusOK, "admin/index.html", adminR)
		//log.Println(adminR)
	}()

	val, b := c.Get("visitnum")
	if !b {
		log.Println("get visit num from midware failed")
		adminR.VisitNum = val.(int64)
	} else {
		adminR.VisitNum = val.(int64)
	}

	menus, err := mongocli.GetMenuListByParent("")
	//log.Println("menus lv1 are ", menus)
	if err != nil {
		res = "get menu list failed"
		adminR.Res = res
		log.Println("get menulv1 failed, err is ", err)
		return
	}

	//获取菜单信息
	for _, menu := range menus {
		menulv1 := &model.MenuLv1{}
		menulv1.SelfMenu = menu
		adminR.Menus = append(adminR.Menus, menulv1)
		childmenus, err := mongocli.GetMenuListByParent(menu.CatId)
		if err != nil {
			res = "get menu list failed"
			adminR.Res = res
			log.Println("get menulv2 ", menu.CatId, " failed, err is ", err)
			menulv1.ChildMenu = []*model.CatMenu{}
			continue
		}
		menulv1.ChildMenu = childmenus
		//log.Println("menus lv1 ChildMenu are ", childmenus)
	}

	//统计总数
	count, err := mongocli.ArticleTotalCount()
	if err != nil {
		adminR.Res = "get article total count failed"
		log.Println("get article total count failed, err is ", err)
		return
	}

	var page_f float64 = float64(count) / 5
	adminR.Total = int(math.Ceil(page_f))
	adminR.Cur = 1

	//获取文章列表
	articles, err := mongocli.GetArticlesByPage(1)
	if err != nil {
		res = "get article list failed"
		adminR.Res = res
		log.Println("get menulv1 failed, err is ", err)
		return
	}

	adminR.Articles = []*model.ArticleR{}
	for _, article := range articles {
		articleR := &model.ArticleR{}
		articleR.Id = article.Id
		articleR.Author = article.Author
		articleR.Cat = article.Cat
		createtm := time.Unix(article.CreateAt, 0)
		articleR.CreateAt = createtm.Format("2006-01-02 15:04:05")
		lasttm := time.Unix(article.LastEdit, 0)
		articleR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
		articleR.LoveNum = article.LoveNum
		articleR.ScanNum = article.ScanNum
		articleR.Subcat = article.Subcat
		articleR.Subtitle = article.Subtitle
		articleR.Title = article.Title

		adminR.Articles = append(adminR.Articles, articleR)
	}
	adminR.Res = res
}

func ArticleEdit(c *gin.Context) {
	subcat := c.Query("subcat")
	cat := c.Query("cat")
	log.Println("subcat is ", subcat)
	log.Println("cat is ", cat)

	articleR := model.ArticleEditR{}
	articleR.Res = model.RENDER_MSG_SUCCESS
	articleR.Cat = cat
	articleR.SubCat = subcat
	defer func() {
		c.HTML(http.StatusOK, "admin/articleedit.html", articleR)
	}()
	menu_, err := mongocli.GetMenuListByParent("")
	if err != nil {
		log.Println("get menu lv1 failed, err is ", err)
		articleR.Res = "get menu failed"
		return
	}
	articleR.Menu_ = menu_

	if subcat == "" && cat == "" {
		return
	}

	submenu_, err := mongocli.GetMenuListByParent(articleR.Cat)
	if err != nil {
		log.Println("get menu lv1 failed, err is ", err)
		articleR.Res = "get menu failed"
		return
	}

	articleR.SubMenu_ = submenu_

	CatEle, err := mongocli.GetMenuById(cat)
	if err != nil {
		log.Println("get menu by id failed, err is ", err)
		articleR.Res = "get menu failed"
		return
	}

	subCatEle, err := mongocli.GetMenuById(subcat)
	if err != nil {
		log.Println("get menu by id failed, err is ", err)
		articleR.Res = "get menu failed"
		return
	}

	articleR.SubCatName = subCatEle.Name
	articleR.CatName = CatEle.Name

}

func SubCatSelect(c *gin.Context) {
	selectR := model.SubCatSelectR{}
	selectR.Res = model.RENDER_MSG_SUCCESS
	subCat := model.SubCatSelectReq{}
	defer func() {
		c.HTML(http.StatusOK, "admin/subcatselect.html", selectR)
	}()
	err := c.BindJSON(&subCat)
	if err != nil {
		log.Println(model.MSG_JSON_UNPACK)
		selectR.Res = model.MSG_JSON_UNPACK
		return
	}
	log.Println("subCat id is ", subCat.CatId)
	menu_, err := mongocli.GetMenuListByParent(subCat.CatId)
	if err != nil {
		log.Println("get menu lv2 failed, err is ", err)
		selectR.Res = "get menu lv2 failed"
		return
	}
	selectR.SubCatMenus_ = menu_
}

func ArticlePub(c *gin.Context) {
	articlePub := model.ArticlePubReq{}
	rsp := model.ArticlePubRsp{}
	rsp.Code = model.SUCCESS_NO
	rsp.Msg = model.MSG_SUCCESS
	err := c.BindJSON(&articlePub)
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
	if err != nil {
		log.Println(model.MSG_JSON_UNPACK)
		rsp.Msg = model.MSG_JSON_UNPACK
		rsp.Code = model.ERR_JSON_UNPACK
		return
	}

	log.Println("articlePub subcat is ", articlePub.Subcat)
	maxindex, err := mongocli.GetSubCatMaxIndex(articlePub.Subcat)
	if err != nil {
		log.Println("get max index failed err is ", err)
	}

	log.Println("max index is ", maxindex)

	log.Println("articlepub is ", articlePub)
	articledb := &model.ArticleInfo{}
	articledb.Id = ksuid.New().String()
	articledb.Author = articlePub.Author
	articledb.Cat = articlePub.Cat

	articledb.CreateAt = time.Now().Local().Unix()
	articledb.LastEdit = time.Now().Local().Unix()
	// log.Println("createat is ", articledb.CreateAt)
	// log.Println("LastEdit is ", articledb.LastEdit)

	articledb.Subcat = articlePub.Subcat
	articledb.Subtitle = articlePub.Subtitle
	articledb.Title = articlePub.Title
	articledb.LoveNum = 500 + rand.Intn(100)
	articledb.ScanNum = 1000 + rand.Intn(100)
	articledb.Index = maxindex + 1
	//存储文章信息
	err = mongocli.SaveArtInfo(articledb)
	if err != nil {
		log.Println(model.MSG_JSON_UNPACK)
		rsp.Msg = model.MSG_SAVE_ARTICLE
		rsp.Code = model.ERR_SAVE_ARTICLE
		return
	}

	contentdb := &model.ArticleContent{}
	contentdb.Id = articledb.Id
	contentdb.Content = articlePub.Content

	err = mongocli.SaveArtContent(contentdb)
	if err != nil {
		log.Println(model.MSG_JSON_UNPACK)
		rsp.Msg = model.MSG_SAVE_ARTICLE
		rsp.Code = model.ERR_SAVE_ARTICLE
		return
	}
}

func Login(c *gin.Context) {
	log.Println("receive admin login ")
	c.HTML(http.StatusOK, "admin/login.html", nil)
}

func LoginSub(c *gin.Context) {
	loginReq := model.LoginSubReq{}
	err := c.BindJSON(&loginReq)
	if err != nil {
		log.Println(model.MSG_JSON_UNPACK)
		rsp := model.LoginSubRsp{}
		rsp.Msg = model.MSG_JSON_UNPACK
		rsp.Code = model.ERR_JSON_UNPACK
		c.JSON(http.StatusOK, rsp)
		return
	}
	log.Println("loginsub is ", loginReq)

	//判断一分中内登录失败次数
	loginfailed, err := mongocli.GetLoginFailed(loginReq.Email)
	if err == nil {
		if loginfailed.Count >= 3 {
			log.Println("get failed count limit per minitue")
			rsp := model.LoginSubRsp{}
			rsp.Msg = model.MSG_LOGIN_LIMIT
			rsp.Code = model.ERR_LOGIN_LIMIT
			c.JSON(http.StatusOK, rsp)
			return
		}
	} else {
		log.Println("get failed count failed")
	}

	//取出数据库存储的密码
	admin, err := mongocli.GetAdminByEmail(loginReq.Email)
	if err != nil {
		log.Println("get admin by email failed")

		rsp := model.LoginSubRsp{}
		rsp.Msg = model.MSG_NO_EMAIL
		rsp.Code = model.ERR_NO_EMAIL
		c.JSON(http.StatusOK, rsp)
		return
	}

	pwdsecret := sha1.Sum([]byte(admin.Pwd + loginReq.Salt))
	pwdsecretStr := fmt.Sprintf("%x", pwdsecret)

	if pwdsecretStr != loginReq.Pwd {
		log.Println("login sub pwd is error")
		//设置失败次数
		if loginfailed == nil {
			loginfailed = &model.LoginFailed{Email: loginReq.Email, Count: 1, CreatedAt: time.Now()}
			mongocli.SaveLoginFailed(loginfailed)
		} else {
			loginfailed.Count++
			log.Println("loginfailed is ", loginfailed)
			err = mongocli.UpdateLoginFailed(loginfailed)
			if err != nil {
				log.Println("update loginfailed count err is ", err)
			}
		}

		rsp := model.LoginSubRsp{}
		rsp.Msg = model.MSG_LOGIN_PWD_ERR
		rsp.Code = model.ERR_PWD_ERROR
		c.JSON(http.StatusOK, rsp)
		return
	}

	//将session保存至数据库
	sid := ksuid.New()
	sessionData := &model.Session{Sid: sid.String(), CreatedAt: time.Now()}
	_, err = mongocli.SaveSession(sessionData)
	if err != nil {
		log.Println("insert session failed, err is ", err)
		rsp := model.LoginSubRsp{}
		rsp.Msg = model.MSG_SAVE_SESSION
		rsp.Code = model.ERR_NO_SESSION
		c.JSON(http.StatusOK, rsp)
		return
	}

	// log.Println("config.TotalCfgData.Cookie.Host is ", config.TotalCfgData.Cookie.Host)
	// log.Println("config.TotalCfgData.Cookie.Alive is ", config.TotalCfgData.Cookie.Alive)
	c.SetCookie(model.CookieSession, sid.String(), config.TotalCfgData.Cookie.Alive, "/",
		config.TotalCfgData.Cookie.Host, false, true)

	rsp := model.LoginSubRsp{}
	rsp.Msg = model.MSG_SUCCESS
	rsp.Code = model.SUCCESS_NO
	rsp.Email = admin.Email

	rsp.Pwd = pwdsecretStr
	c.JSON(http.StatusOK, rsp)
}

func ArticleSearch(c *gin.Context) {
	searchReq := model.SearchArticleReq{}
	articleListR := model.ArticleListR{}
	articleListR.Res = model.RENDER_MSG_SUCCESS
	articleListR.Articles = []*model.ArticleR{}
	defer func() {
		c.HTML(http.StatusOK, "admin/articlelist.html", articleListR)
	}()
	err := c.BindJSON(&searchReq)
	if err != nil {
		log.Println(model.MSG_JSON_UNPACK)
		articleListR.Res = model.MSG_JSON_UNPACK
		return
	}

	log.Println("search article condition  is ", searchReq)

	if searchReq.Page <= 0 {
		log.Println(model.MSG_INVALID_PARAM)
		articleListR.Res = model.MSG_INVALID_PARAM
		return
	}

	artileList, total, err := mongocli.SearchArticle(&searchReq)
	if err != nil {
		log.Println("search article list failed err is ", err)
		articleListR.Res = "search article list failed"
		return
	}

	for _, article := range artileList {
		articleR := &model.ArticleR{}
		articleR.Id = article.Id
		articleR.Author = article.Author
		articleR.Cat = article.Cat
		createtm := time.Unix(article.CreateAt, 0)
		articleR.CreateAt = createtm.Format("2006-01-02 15:04:05")
		lasttm := time.Unix(article.LastEdit, 0)
		articleR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
		articleR.LoveNum = article.LoveNum
		articleR.ScanNum = article.ScanNum
		articleR.Subcat = article.Subcat
		articleR.Subtitle = article.Subtitle
		articleR.Title = article.Title

		articleListR.Articles = append(articleListR.Articles, articleR)
	}
	articleListR.Res = model.RENDER_MSG_SUCCESS

	var page_f float64 = float64(total) / 5

	articleListR.Cur = searchReq.Page
	articleListR.Total = int(math.Ceil(page_f))
	log.Println("cur is ", articleListR.Cur)
	log.Println("total is ", articleListR.Total)
	log.Println("page_f is ", page_f)
}

func DelArticle(c *gin.Context) {
	delArticleReq := &model.DelArticleReq{}
	delRsp := model.DelArticleRsp{}
	delRsp.Code = model.SUCCESS_NO
	delRsp.Msg = model.MSG_SUCCESS
	defer func() {
		c.JSON(http.StatusOK, delRsp)
	}()

	err := c.BindJSON(delArticleReq)
	if err != nil {
		log.Println("json parse failed, err is ", err)
		delRsp.Code = model.ERR_JSON_UNPACK
		delRsp.Msg = model.MSG_JSON_UNPACK
		return
	}
	log.Println("del article req is ", delArticleReq)

	err = mongocli.DelArticle(delArticleReq.Title)
	if err != nil {
		log.Println("delete article failed")
		delRsp.Code = model.ERR_DEL_ARTICLE
		delRsp.Msg = model.MSG_DEL_ARTICLE
		return
	}
}

func ModifyArticle(c *gin.Context) {
	id := c.Query("id")
	modify := model.ArticleModifyR{}
	modify.Res = model.MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "admin/articlemodify.html", modify)
	}()

	article, err := mongocli.GetArticleId(id)
	if err != nil {
		log.Println("get article failed, err is ", err)
		modify.Res = "get article failed"
		return
	}

	menu_, err := mongocli.GetMenuListByParent("")
	if err != nil {
		log.Println("get menu lv1 failed, err is ", err)
		modify.Res = "get menu failed"
		return
	}
	modify.Menu_ = menu_

	submenu_, err := mongocli.GetMenuListByParent(article.Cat)
	if err != nil {
		log.Println("get menu lv1 failed, err is ", err)
		modify.Res = "get menu failed"
		return
	}

	modify.SubMenu_ = submenu_
	log.Println("article cat is ", article.Cat)
	CatEle, err := mongocli.GetMenuByCat(article.Cat)
	if err != nil {
		log.Println("get menu by id failed, err is ", err)
		modify.Res = "get menu failed"
		return
	}

	log.Println("article subcat is ", article.Subcat)
	subCatEle, err := mongocli.GetMenuByCat(article.Subcat)
	if err != nil {
		log.Println("get menu by id failed, err is ", err)
		modify.Res = "get menu failed"
		return
	}

	modify.SubCatName = subCatEle.Name
	modify.CatName = CatEle.Name

	articleR := &model.ArticleR{}
	articleR.Id = article.ArticleInfo.Id
	articleR.Author = article.Author
	articleR.Cat = article.Cat
	articleR.Content = article.Content
	createtm := time.Unix(article.CreateAt, 0)
	articleR.CreateAt = createtm.Format("2006-01-02 15:04:05")
	lasttm := time.Unix(article.LastEdit, 0)
	articleR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
	articleR.LoveNum = article.LoveNum
	articleR.ScanNum = article.ScanNum
	articleR.Subcat = article.Subcat
	articleR.Subtitle = article.Subtitle
	articleR.Title = article.Title

	modify.Article = articleR
	log.Println("modify article is ", articleR)
}

func UpdateArticle(c *gin.Context) {
	updatearticle := &model.UpdateArticleReq{}
	rsp := model.UpdateArticleRsp{}
	rsp.Code = model.SUCCESS_NO
	rsp.Msg = model.MSG_SUCCESS
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	err := c.BindJSON(updatearticle)
	if err != nil {
		log.Println("json parse failed, err is ", err)
		rsp.Code = model.ERR_JSON_UNPACK
		rsp.Msg = model.MSG_JSON_UNPACK
		return
	}
	log.Println("update article req is ", updatearticle)

	err = mongocli.UpdateArticle(updatearticle)
	if err != nil {
		log.Println("update article failed")
		rsp.Code = model.ERR_UPDATE_ARTICLE
		rsp.Msg = model.MSG_UPDATE_ARTICLE
		return
	}
}
