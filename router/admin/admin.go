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
	article_edit := model.ArticleEditReq{}
	c.BindJSON(&article_edit)
	log.Printf("article_edit data is %v", &article_edit)
}
