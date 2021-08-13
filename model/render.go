package model

//管理主界面
type AdminIndexR struct {
	Menus []*MenuLv1
	Res   string
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

/*
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
*/
