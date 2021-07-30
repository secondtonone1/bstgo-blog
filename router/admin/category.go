package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Category(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/articlecateg.html", nil)
}

func Sort(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/articlesort.html", nil)
}

func SortSave(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/articlecateg.html", nil)
}

func IndexList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/indexlist.html", nil)
}
