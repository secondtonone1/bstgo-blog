package admin

import (
	"bstgo-blog/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", nil)
}

func ArticleEdit(c *gin.Context) {
	subcat := c.Query("subcat")
	cat := c.Query("cat")
	log.Println("subcat is ", subcat)
	log.Println("cat is ", cat)
	c.HTML(http.StatusOK, "admin/articleedit.html", nil)
}

func ArticlePub(c *gin.Context) {
	articlePub := model.ArticlePubReq{}
	c.BindJSON(&articlePub)
	log.Println("articlepub is ", articlePub)
	c.JSON(http.StatusOK, articlePub)
}
