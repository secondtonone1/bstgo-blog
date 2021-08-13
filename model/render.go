package model

//管理主界面
type AdminIndexR struct {
	Menu_ *Menu
	Res   string
}

//文章编辑界面渲染
type ArticleEditR struct {
	Menu_ *Menu
	Res   string
}

//子类别渲染
type SubCatSelectR struct {
	SubCatMenus_ []*SubCatMenu
	Res          string
}
