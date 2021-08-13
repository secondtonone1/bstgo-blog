package admin

import (
	"bstgo-blog/config"
	"bstgo-blog/model"
	mongocli "bstgo-blog/mongo"
	"fmt"
	"log"
	"net/http"
	"time"

	"crypto/sha1"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func Admin(c *gin.Context) {
	var res string = ""
	adminR := model.AdminIndexR{}
	defer func() {
		c.HTML(http.StatusOK, "admin/index.html", adminR)
	}()
	menus, err := mongocli.GetMenuList()

	if err != nil {
		res = "get menu list failed"
		adminR.Res = res
		return
	}

	adminR.Menu_ = menus
	adminR.Res = res
}

func ArticleEdit(c *gin.Context) {
	subcat := c.Query("subcat")
	cat := c.Query("cat")
	log.Println("subcat is ", subcat)
	log.Println("cat is ", cat)
	articleR := model.ArticleEditR{}
	articleR.Res = model.RENDER_MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "admin/articleedit.html", articleR)
	}()
	menu, err := mongocli.GetMenuList()
	if err != nil {
		log.Println("get menu failed")
		articleR.Res = "get menu failed"
		return
	}

	articleR.Menu_ = menu
}

func SubCatSelect(c *gin.Context) {
	subCat := model.SubCatSelectReq{}
	c.BindJSON(&subCat)
	log.Println("subCat id is ", subCat.CatId)
	selectR := model.SubCatSelectR{}
	selectR.Res = model.RENDER_MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "admin/subcatselect", selectR)
	}()
	menu, err := mongocli.GetSubCatSelect(subCat.CatId)
	if err != nil {
		log.Println("get subcat select failed, err is ", err)
		return
	}

	selectR.SubCatMenus_ = menu.CatMenus_[0].SubCatMenus_
	// for _, val := range selectR.SubCatMenus_ {
	// 	log.Println("subcat name is ", val.Name)
	// }
	for _, val := range menu.CatMenus_ {
		log.Println(val.Name)
	}
}

func ArticlePub(c *gin.Context) {
	articlePub := model.ArticlePubReq{}
	c.BindJSON(&articlePub)
	log.Println("articlepub is ", articlePub)
	c.JSON(http.StatusOK, articlePub)
}

func Login(c *gin.Context) {
	log.Println("receive admin login ")
	c.HTML(http.StatusOK, "admin/login.html", nil)
}

func LoginSub(c *gin.Context) {
	loginReq := model.LoginSubReq{}
	c.BindJSON(&loginReq)
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
