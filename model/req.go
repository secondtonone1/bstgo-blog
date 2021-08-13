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
	Cat     string `json: "cat"`
	Title   string `json: "title"`
	Content string `json: "content"`
	Subcat  string `json: "subcat"`
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

type SubMenu struct {
	SubCatId string `json:"subcatid" bson:"subcatid"`
	Name     string `json:"name" bson:"name"`
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
