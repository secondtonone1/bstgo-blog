package admin

import (
	"bstgo-blog/model"
	mongocli "bstgo-blog/mongo"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func Category(c *gin.Context) {
	catReq := &model.AdminCatReq{}
	adminCat := &model.AdminCatR{}
	adminCat.Articles = []*model.ArticleR{}
	adminCat.Res = model.MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "admin/articlecateg.html", adminCat)
	}()

	err := c.BindJSON(catReq)
	if err != nil {
		log.Println("json parse failed, err is ", err)
		adminCat.Res = "json parse failed"
		return
	}
	log.Println("admin cat is ", catReq)

	articles, err := mongocli.SearchArticleBySubCat(catReq.Category)
	if err != nil {
		log.Println("get articles failed")
		adminCat.Res = "get articles failed"
		return
	}

	log.Println("articles are ", articles)
	for _, article := range articles {
		//log.Println("article is ", article)
		articleR := &model.ArticleR{}
		articleR.Id = article.Id
		articleR.Author = article.Author
		articleR.Cat = article.Cat
		articleR.Content = ""
		createtm := time.Unix(article.CreateAt, 0)
		articleR.CreateAt = createtm.Format("2006-01-02 15:04:05")
		lasttm := time.Unix(article.LastEdit, 0)
		articleR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
		articleR.LoveNum = article.LoveNum
		articleR.ScanNum = article.ScanNum
		articleR.Subcat = article.Subcat
		articleR.Subtitle = article.Subtitle
		articleR.Title = article.Title
		adminCat.Articles = append(adminCat.Articles, articleR)
	}

}

func Sort(c *gin.Context) {
	catReq := &model.AdminCatReq{}
	adminCat := &model.AdminCatR{}
	adminCat.Articles = []*model.ArticleR{}
	adminCat.Res = model.MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "admin/articlesort.html", adminCat)
	}()

	err := c.BindJSON(catReq)
	if err != nil {
		log.Println("json parse failed, err is ", err)
		adminCat.Res = "json parse failed"
		return
	}

	log.Println("admin cat is ", catReq)

	articles, err := mongocli.SearchArticleBySubCat(catReq.Category)
	if err != nil {
		log.Println("get articles failed")
		adminCat.Res = "get articles failed"
		return
	}

	for _, article := range articles {
		articleR := &model.ArticleR{}
		articleR.Id = article.Id
		articleR.Author = article.Author
		articleR.Cat = article.Cat
		articleR.Content = ""
		createtm := time.Unix(article.CreateAt, 0)
		articleR.CreateAt = createtm.Format("2006-01-02 15:04:05")
		lasttm := time.Unix(article.LastEdit, 0)
		articleR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
		articleR.LoveNum = article.LoveNum
		articleR.ScanNum = article.ScanNum
		articleR.Subcat = article.Subcat
		articleR.Subtitle = article.Subtitle
		articleR.Title = article.Title
		adminCat.Articles = append(adminCat.Articles, articleR)
	}

}

func SortSave(c *gin.Context) {
	sortArtReq := &model.ArticleSortReq{}
	adminCat := &model.AdminCatR{}
	adminCat.Articles = []*model.ArticleR{}
	adminCat.Res = model.MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "admin/articlecateg.html", adminCat)
	}()

	err := c.BindJSON(sortArtReq)
	if err != nil {
		log.Println("json parse failed, err is ", err)
		adminCat.Res = "json parse failed"
		return
	}

	log.Println("sortArtReq cat is ", sortArtReq)

	err = mongocli.UpdateArticleSort(sortArtReq)
	if err != nil {
		log.Println("update article sort failed, err is ", err)
		adminCat.Res = "update article sort failed"
		return
	}

	articles, err := mongocli.SearchArticleBySubCat(sortArtReq.SubCat)

	if err != nil {
		log.Println("get articles failed")
		adminCat.Res = "get articles failed"
		return
	}

	for _, article := range articles {
		articleR := &model.ArticleR{}
		articleR.Id = article.Id
		articleR.Author = article.Author
		articleR.Cat = article.Cat
		articleR.Content = ""
		createtm := time.Unix(article.CreateAt, 0)
		articleR.CreateAt = createtm.Format("2006-01-02 15:04:05")
		lasttm := time.Unix(article.LastEdit, 0)
		articleR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
		articleR.LoveNum = article.LoveNum
		articleR.ScanNum = article.ScanNum
		articleR.Subcat = article.Subcat
		articleR.Subtitle = article.Subtitle
		articleR.Title = article.Title
		adminCat.Articles = append(adminCat.Articles, articleR)
	}
}

func IndexList(c *gin.Context) {

	var res string = model.RENDER_MSG_SUCCESS
	adminR := model.AdminIndexR{}
	adminR.Cur = 1
	adminR.Total = 1
	adminR.Res = res
	defer func() {
		c.HTML(http.StatusOK, "admin/indexlist.html", adminR)
		//log.Println(adminR)
	}()

	req := &model.SearchArticleReq{}
	err := c.BindJSON(req)
	if err != nil {
		res = "json parse failed "
		adminR.Res = res
		log.Println("get menulv1 failed, err is ", err)
		return
	}

	if req.Page <= 0 {
		res = "page less than 0"
		adminR.Res = res
		log.Println(res)
		return
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

	//获取文章列表
	articles, total, err := mongocli.SearchArticle(req)
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
		articleR.Content = ""
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

	var page_f float64 = float64(total) / 5
	adminR.Total = int(math.Ceil(page_f))
	adminR.Cur = req.Page
}

//创建分类
func CreateCtg(c *gin.Context) {
	create_ctg := model.CreateCtgReq{}
	menuR := &model.MenuLv1R{}
	defer func() {
		c.HTML(http.StatusOK, "admin/ctgele.html", menuR)
	}()

	err := c.BindJSON(&create_ctg)
	if err != nil {
		menuR.Msg = model.MSG_JSON_UNPACK
		return
	}
	log.Printf("%v", &create_ctg)
	id := ksuid.New()

	catmenu := &model.CatMenu{CatId: id.String(),
		Name:   create_ctg.Category,
		Parent: "",
		Index:  create_ctg.Index}
	_, err = mongocli.SaveMenu(catmenu)

	menuR.CatId = id.String()
	menuR.CatName = create_ctg.Category
	menuR.Msg = model.RENDER_MSG_SUCCESS

	if err != nil {
		menuR.Msg = "menu lv1 insert failed!"
		log.Println(menuR.Msg)
	}
}

//创建子分类
func CreateSubCtg(c *gin.Context) {
	menulv2 := &model.MenuLv2R{}
	defer func() {
		c.HTML(http.StatusOK, "admin/subctgele.html", menulv2)
	}()
	create_subctg := model.CreateSubCtgReq{}
	err := c.BindJSON(&create_subctg)
	if err != nil {
		menulv2.Msg = model.MSG_JSON_UNPACK
		return
	}
	log.Printf("%v", &create_subctg)
	id := ksuid.New()

	menulv2.Msg = model.RENDER_MSG_SUCCESS
	menulv2.SubCatId = id.String()
	menulv2.SubCatName = create_subctg.SubCategory

	catmenu := &model.CatMenu{CatId: id.String(),
		Name:   create_subctg.SubCategory,
		Parent: create_subctg.ParentId,
		Index:  create_subctg.Index}
	_, err = mongocli.SaveMenu(catmenu)

	if err != nil {
		menulv2.Msg = "save menu lv2 failed"
	}

}

//排序子菜单
func SortMenu(c *gin.Context) {

	rsp := model.SortMenuRsp{}
	rsp.Code = model.SUCCESS_NO
	rsp.Msg = model.MSG_SUCCESS
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	sortmenu := model.SortMenuReq{}
	err := c.BindJSON(&sortmenu)
	if err != nil {
		log.Println("parase json failed, err is ", err)
		rsp.Code = model.SUCCESS_NO
		rsp.Msg = model.MSG_SUCCESS
		return
	}

	for _, menu := range sortmenu.Menulist {
		log.Println("menu catid is ", menu.CatId)
		log.Println("menu index is ", menu.Index)
	}
	err = mongocli.SaveMenuList(sortmenu.Menulist)
	if err != nil {
		log.Println(model.MSG_SAVE_MENUS)
		rsp.Code = model.ERR_SAVE_MENUS
		rsp.Msg = model.MSG_SAVE_MENUS
		return
	}
}
