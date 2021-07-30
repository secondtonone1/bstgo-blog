package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	//c.String(http.StatusOK, "Hello World")
	c.HTML(http.StatusOK, "home/index.html", nil)
}

func Category(c *gin.Context) {
	c.HTML(http.StatusOK, "home/categary.html", nil)
}
