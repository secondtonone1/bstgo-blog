package model

import "html/template"

//管理主界面
type AdminIndexR struct {
	Menus    []*MenuLv1
	Res      string
	Articles []*ArticleR
	Total    int
	Cur      int
	VisitNum int64
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
	HomeArticleR
	RelRecommendR
	HotListR
	NewCommentR
	NavCatListR
}

//文章页面文章渲染结构
type HomeArticleR struct {
	Id         string `bson:"id"`
	Cat        string `bson: "cat"`
	Title      string `bson: "title"`
	Content    template.HTML
	Subcat     string `bson: "subcat"`
	Subtitle   string `bson: "subtitle"`
	ScanNum    int    `bson:"scannum"`
	LoveNum    int    `bson:"lovenum`
	CreateAt   string `bson:"createdAt"`
	LastEdit   string `bson:"lastedit"`
	Author     string `bson:"author"`
	Index      int    `bson:"index"`
	CommentNum int
	Comments   []*CommentR
	VisitNum   int64
}

//顶部分类信息
type NavCatR struct {
	CatId string
	Name  string
	Index int
}

//顶部菜单列表
type NavCatListR struct {
	NavCatList []*NavCatR
	ActiveId   string
}

//相关推荐
type RelRecommendR struct {
	RecommendList []*HomeArticleR
}

//热门列表
type HotListR struct {
	HotList []*HomeArticleR
}

//最新评论列表
type NewCommentR struct {
	CommentList []*CommentR
}

//文章评论
type CommentR struct {
	Id       string        `bson:"id"`
	UserName string        `bson:"username"`
	Time     string        `bson:"comtime"`
	Content  template.HTML `bson:"content"`
	LoveNum  int           `bson:"lovenum"`
	Parent   string        `bson:"parent"`
	HeadIcon string        `bson:"headicon"`
	ReplyNum int
	Replys   []*ReplyR
	ArtId    string
	ArtTitle string
}

type ReplyR struct {
	Id       string        `bson:"id"`
	UserName string        `bson:"username"`
	Time     string        `bson:"comtime"`
	Content  template.HTML `bson:"content"`
	LoveNum  int           `bson:"lovenum"`
	HeadIcon string        `bson:"headicon"`
	Parent   string        `bson:"parent"`
}

type IndexArticlesR struct {
	IndexArticleList []*HomeArticleR
}

type HomeIndexR struct {
	VisitNum int64
	NavCatListR
	HotListR
	NewCommentR
	IndexArticlesR
	CurPage   int
	TotalPage int
	NextPage  int
}

type HomeCategoryR struct {
	VisitNum     int64
	CategoryName string
	NavCatListR
	HomeArticleR
	LeftCatMenusR
	HotListR
	NewCommentR
}

//左侧分类菜单
type LeftCatMenusR struct {
	LeftCatList []*LeftCatR
}

type LeftCatR struct {
	CatId      string
	Name       string
	SubArticle []*ArtInfoR
}

type CommentRsp struct {
	Res string
	CommentR
}

type CommentReplyR struct {
	Res string
	ReplyR
}

type ArtInfoR struct {
	ArtId       string
	ArtSubTitle string
}

type SubCatArtInfoR struct {
	SubCatArtInfos []*ArtInfoR
	Msg            string
}

type ArticleDetailR struct {
	HomeArticleR
	Msg string
}

type ArticleDetailsR struct {
	Msg       string
	CurPage   int
	TotalPage int
	NextPage  int
	PrevPage  int
	IndexArticlesR
}
