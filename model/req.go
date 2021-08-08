package model

type CreateCtgReq struct {
	Category string `json: "category" `
}

type CreateSubCtgReq struct {
	SubCategory string ` json: "subcategory"`
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
