package model

import "html/template"

//管理主界面
type AdminIndexR struct {
	Menus    []*MenuLv1
	Res      string
	Articles []*ArticleR
	Total    int
	Cur      int
}

//一级目录
type MenuLv1 struct {
	SelfMenu  *CatMenu
	ChildMenu []*CatMenu
}

//一级目录模板
type MenuLv1R struct {
	CatName string
	CatId   string
	Msg     string
}

//二级目录模板
type MenuLv2R struct {
	SubCatName string
	SubCatId   string
	Msg        string
}

//文章编辑界面渲染
type ArticleEditR struct {
	Menu_      []*CatMenu
	SubMenu_   []*CatMenu
	Res        string
	Cat        string
	SubCat     string
	CatName    string
	SubCatName string
}

//文章编辑页面子分类渲染
type SubCatSelectR struct {
	SubCatMenus_ []*CatMenu
	Res          string
}

//文章结构
type ArticleR struct {
	Id       string `bson:"id"`
	Cat      string `bson: "cat"`
	Title    string `bson: "title"`
	Content  string `bson: "content"`
	Subcat   string `bson: "subcat"`
	Subtitle string `bson: "subtitle"`
	ScanNum  int    `bson:"scannum"`
	LoveNum  int    `bson:"lovenum`
	CreateAt string
	LastEdit string
	Author   string `bson:"author"`
}

//admin搜索界面渲染文章列表
type ArticleListR struct {
	Articles []*ArticleR
	Res      string
	Cur      int
	Total    int
}

//点击三级标题渲染文章列表
type AdminCatR struct {
	Res      string
	Articles []*ArticleR
}

//文章修改
type ArticleModifyR struct {
	Res     string
	Article *ArticleR

	Menu_    []*CatMenu
	SubMenu_ []*CatMenu

	Cat        string
	SubCat     string
	CatName    string
	SubCatName string
}

//草稿箱
type DraftBoxR struct {
	Res      string
	Articles []*ArticleR
	Total    int
	Cur      int
}

//文章页模板渲染
type ArticlePageR struct {
	Id       string `bson:"id"`
	Cat      string `bson: "cat"`
	Title    string `bson: "title"`
	Content  template.HTML
	Subcat   string `bson: "subcat"`
	Subtitle string `bson: "subtitle"`
	ScanNum  int    `bson:"scannum"`
	LoveNum  int    `bson:"lovenum`
	CreateAt string `bson:"createdAt"`
	LastEdit string `bson:"lastedit"`
	Author   string `bson:"author"`
	Index    int    `bson:"index"`
}
