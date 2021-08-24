package home

import (
	"bstgo-blog/model"
	mongocli "bstgo-blog/mongo"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	//c.String(http.StatusOK, "Hello World")
	c.HTML(http.StatusOK, "home/index.html", nil)
}

func Category(c *gin.Context) {
	c.HTML(http.StatusOK, "home/categary.html", nil)
}

func ArticlePage(c *gin.Context) {
	id := c.Query("id")
	log.Println("id is ", id)
	if id == "" {
		c.HTML(http.StatusOK, "home/errorpage.html", "invalid page request , id is null, after 2 seconds return to home")
		return
	}

	article, err := mongocli.GetArticleId(id)

	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get article failed, after 2 seconds return to home")
		return
	}

	articleR := &model.ArticlePageR{}
	articleR.Author = article.Author
	articleR.Cat = article.Cat
	articleR.Content = template.HTML(article.Content)
	createtm := time.Unix(article.CreateAt, 0)
	articleR.CreateAt = createtm.Format("2006-01-02 15:04:05")

	lasttm := time.Unix(article.LastEdit, 0)
	articleR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
	articleR.Id = article.Id
	articleR.Index = article.Index
	articleR.LoveNum = article.LoveNum
	articleR.ScanNum = article.ScanNum
	articleR.Subcat = article.Subcat
	articleR.Subtitle = article.Subtitle
	articleR.Title = article.Title

	c.HTML(http.StatusOK, "home/articlepage.html", articleR)
}
