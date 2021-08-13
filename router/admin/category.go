package admin

import (
	"bstgo-blog/model"
	mongocli "bstgo-blog/mongo"
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

//创建分类
func CreateCtg(c *gin.Context) {
	create_ctg := model.CreateCtgReq{}
	c.BindJSON(&create_ctg)
	log.Printf("%v", &create_ctg)
	id := ksuid.New()

	menusExist := true
	menus, err := mongocli.GetMenuList()
	if err != nil {
		log.Println("get menu list failed")
		menus = &model.Menu{}
		menus.CatMenus_ = make([]*model.CatMenu, 0)
		menusExist = false
	}

	for _, catmenu := range menus.CatMenus_ {
		if catmenu.Name == create_ctg.Category {
			log.Println("catemenu name exists")
			c.HTML(http.StatusOK, "admin/ctgele.html", gin.H{
				"catename": create_ctg.Category,
				"cateid":   id.String(),
				"msg":      "catemenu name exists",
			})
			return
		}
	}

	catmenu := &model.CatMenu{Name: create_ctg.Category, CatId: id.String()}
	menus.CatMenus_ = append(menus.CatMenus_, catmenu)

	if !menusExist {
		err = mongocli.SaveMenuList(menus)
		if err != nil {
			log.Println("save menu failed")
			c.HTML(http.StatusOK, "admin/ctgele.html", gin.H{
				"catename": create_ctg.Category,
				"cateid":   id.String(),
				"msg":      "save menu failed",
			})
			return
		}
	} else {
		err = mongocli.UpdateMenuList(menus)
		if err != nil {
			log.Println("update menu failed")
			c.HTML(http.StatusOK, "admin/ctgele.html", gin.H{
				"catename": create_ctg.Category,
				"cateid":   id.String(),
				"msg":      "update menu failed",
			})
			return
		}
	}

	c.HTML(http.StatusOK, "admin/ctgele.html", gin.H{
		"catename": create_ctg.Category,
		"cateid":   id.String(),
		"msg":      model.RENDER_MSG_SUCCESS,
	})
}

//创建子分类
func CreateSubCtg(c *gin.Context) {
	create_subctg := model.CreateSubCtgReq{}
	c.BindJSON(&create_subctg)
	log.Printf("%v", &create_subctg)
	id := ksuid.New()

	res := model.RENDER_MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "admin/subctgele.html", gin.H{
			"catename": create_subctg.SubCategory,
			"cateid":   id.String(),
			"msg":      res,
		})
	}()

	menus, err := mongocli.GetMenuList()
	if err != nil {
		log.Println("get menu list failed")
		res = "no-parent category"
		return
	}

	var catMenuTemp *model.CatMenu = nil
	for _, catmenu := range menus.CatMenus_ {
		if catmenu.CatId == create_subctg.ParentId {
			catMenuTemp = catmenu
			break
		}
	}

	if catMenuTemp == nil {
		log.Println("no-parent category")
		res = "no-parent category"
		return
	}

	subcatMenu := &model.SubCatMenu{SubCatId: id.String(), Name: create_subctg.SubCategory}

	catMenuTemp.SubCatMenus_ = append(catMenuTemp.SubCatMenus_, subcatMenu)

	err = mongocli.UpdateMenuList(menus)
	if err != nil {
		log.Println("save menu failed")
		c.HTML(http.StatusOK, "admin/ctgele.html", gin.H{
			"catename": create_subctg.SubCategory,
			"cateid":   id.String(),
			"msg":      "save menu failed",
		})
		return
	}

}

//排序子菜单
func SortMenu(c *gin.Context) {

	rsp := model.SortMenuRsp{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	sortmenu := model.SortMenuReq{}
	c.BindJSON(&sortmenu)
	log.Printf("%v", &sortmenu)
	err := mongocli.UpdateSortMenu(&sortmenu)

	if err != nil {
		log.Println("update sort menu failed, err is ", err)
		rsp.Code = model.ERR_SORT_MENU
		rsp.Msg = model.MSG_SORT_MENU
		return
	}

	rsp.Code = model.SUCCESS_NO
	rsp.Msg = model.MSG_SUCCESS

}
