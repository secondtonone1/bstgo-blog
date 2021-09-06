package model

type CreateCtgReq struct {
	Category string `json: "category" `
	Index    int    `json: "index"`
}

type CreateSubCtgReq struct {
	ParentId    string `json:"parentid"`
	SubCategory string `json:"subcategory"`
	Index       int    `json:"index"`
}

type ArticleParamReq struct {
	Cat    string `json: "cat"`
	SubCat string `json: "subcat"`
}

type ArticlePubReq struct {
	Cat      string `json: "cat"`
	Title    string `json: "title"`
	Content  string `json: "content"`
	Subcat   string `json: "subcat"`
	Subtitle string `json: "subtitle"`
	Author   string `json: "author"`
}

type ArticlePubRsp struct {
	BaseRsp
}

type LoginSubReq struct {
	Pwd   string `json: "pwd"`
	Email string `json: "email"`
	Salt  string `json: "salt"`
}

type BaseRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LoginSubRsp struct {
	BaseRsp
	Pwd   string `json: "pwd"`
	Email string `json: "email"`
}

type SortMenuReq struct {
	Menulist []*CatMenu `json:"menu"`
}

type SortMenuRsp struct {
	BaseRsp
}

type SubCatSelectReq struct {
	CatId string `json:"catid" bson:"catid"`
}

type SearchArticleReq struct {
	Year     string `json:"year"`
	Month    string `json:"month"`
	Cat      string `json:"cat"`
	Keywords string `json:"keywords"`
	Page     int    `json:"page"`
}

type AdminCatReq struct {
	Category string `json:"category"`
}

type ArticleSort struct {
	Id    string `json:"articleid"`
	Title string `json:"title"`
	Index int    `json:"index"`
}

type ArticleSortReq struct {
	ArticleList []*ArticleSort `json:"sortlist"`
	SubCat      string         `json:"subcat"`
}

type DelArticleReq struct {
	Title string `json:"title"`
}

type DelArticleRsp struct {
	BaseRsp
}

type TotalArticleReq struct {
	CurPage int `json:"curpage"`
}

type UpdateArticleReq struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Id       string `json:"id"`
	Cat      string `json:"cat"`
	SubCat   string `json:"subcat"`
	LastEdit int    `json:"lastedit"`
	Author   string `json:"author"`
	Content  string `json:"content"`
}

type UpdateArticleRsp struct {
	BaseRsp
}

type DraftBoxReq struct {
	CurPage int `json:"curpage"`
}

type AddLoveNumReq struct {
	Id string `json:"id"`
}

type AddLoveNumRsp struct {
	BaseRsp
}

type CommentReq struct {
	UserName string `json:"username"`
	HeadIcon string `json:"headicon"`
	Content  string `json:"content"`
	Parent   string `json:"parent"`
	ArtId    string `json:"artid"`
}

type ComLoveReq struct {
	Id string `json:"id"`
}

type ComLoveRsp struct {
	BaseRsp
}

type ComReplyReq struct {
	Parent   string `json:"parent"`
	UserName string `json:"username"`
	HeadIcon string `json:"headicon"`
	Content  string `json:"content"`
	ArtId    string `json:"artid"`
}

type RplLoveReq struct {
	Id string `json:"id"`
}

type RplLoveRsp struct {
	BaseRsp
}

type SubCatArtInfoReq struct {
	Cat    string `json:"cat"`
	SubCat string `json:"subcat"`
}
