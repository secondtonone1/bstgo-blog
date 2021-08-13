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
