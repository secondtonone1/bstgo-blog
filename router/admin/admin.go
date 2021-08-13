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
	var res string = model.RENDER_MSG_SUCCESS
	adminR := model.AdminIndexR{}
	defer func() {
		c.HTML(http.StatusOK, "admin/index.html", adminR)
		//log.Println(adminR)
	}()
	menus, err := mongocli.GetMenuListByParent("")
	//log.Println("menus lv1 are ", menus)
	if err != nil {
		res = "get menu list failed"
		adminR.Res = res
		log.Println("get menulv1 failed, err is ", err)
		return
	}

	for _, menu := range menus {
		menulv1 := &model.MenuLv1{}
		menulv1.SelfMenu = menu
		adminR.Menus = append(adminR.Menus, menulv1)
		childmenus, err := mongocli.GetMenuListByParent(menu.CatId)
		if err != nil {
			res = "get menu lv2 failed"
			adminR.Res = res
			log.Println("get menulv2 ", menu.CatId, " failed, err is ", err)
			menulv1.ChildMenu = []*model.CatMenu{}
			continue
		}
		menulv1.ChildMenu = childmenus
		//log.Println("menus lv1 ChildMenu are ", childmenus)
	}

	adminR.Res = res
}

func ArticleEdit(c *gin.Context) {
	subcat := c.Query("subcat")
	cat := c.Query("cat")
	log.Println("subcat is ", subcat)
	log.Println("cat is ", cat)
}

func SubCatSelect(c *gin.Context) {
	subCat := model.SubCatSelectReq{}
	err := c.BindJSON(&subCat)
	if err != nil {
		log.Println(model.MSG_JSON_UNPACK)
		return
	}
	log.Println("subCat id is ", subCat.CatId)

}

func ArticlePub(c *gin.Context) {
	articlePub := model.ArticlePubReq{}
	err := c.BindJSON(&articlePub)
	defer func() {
		c.JSON(http.StatusOK, articlePub)
	}()
	if err != nil {
		log.Println(model.MSG_JSON_UNPACK)
		return
	}
	log.Println("articlepub is ", articlePub)

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
