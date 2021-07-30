package main

import (
	"bstgo-blog/router/home"
	"net/http"

	"bstgo-blog/router/admin"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")
	router.StaticFS("/static", http.Dir("./public"))
	router.GET("/home", home.Home)
	router.GET("/category", home.Category)
	router.GET("/admin", admin.Admin)
	router.POST("/admin/category", admin.Category)
	router.POST("/admin/sort", admin.Sort)
	router.POST("/admin/sortsave", admin.SortSave)
	router.POST("/admin/index", admin.IndexList)
	router.Run(":8080")
}
