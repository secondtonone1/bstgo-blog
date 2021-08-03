package model

type CreateCtgReq struct {
	Category string `json: "category" `
}

type CreateSubCtgReq struct {
	SubCategory string ` json: "subcategory"`
}

type ArticleEditReq struct {
	Cat    string `json: "cat"`
	SubCat string `json: "subcat"`
}
