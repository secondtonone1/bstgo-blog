package admin

import (
	"bstgo-blog/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
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

func CreateCtg(c *gin.Context) {
	create_ctg := model.CreateCtgReq{}
	c.BindJSON(&create_ctg)
	log.Printf("%v", &create_ctg)
	id := ksuid.New()

	c.HTML(http.StatusOK, "admin/ctgele.html", gin.H{
		"catename": create_ctg.Category,
		"cateid":   id.String(),
	})
}

func CreateSubCtg(c *gin.Context) {
	create_subctg := model.CreateSubCtgReq{}
	c.BindJSON(&create_subctg)
	log.Printf("%v", &create_subctg)
	id := ksuid.New()
	c.HTML(http.StatusOK, "admin/subctgele.html", gin.H{
		"catename": create_subctg.SubCategory,
		"cateid":   id.String(),
	})
}
