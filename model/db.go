package model

import (
	"time"
)

type Session struct {
	CreatedAt time.Time `json:"_" bson:"createdAt"`
	Sid       string    `json:"sid" bson:"sid"`
}

type Admin struct {
	Email string `json: "email" bson:"email"`
	Pwd   string `json: "pwd" bson:"pwd"`
}

type LoginFailed struct {
	CreatedAt time.Time `json:"_" bson:"createdAt"`
	Email     string    `json: "email" bson:"email"`
	Count     int       `json: "count" bson:"count"`
}

//分类目录
type CatMenu struct {
	CatId  string `bson:"catid" json:"catid"`
	Name   string `bson:"name" json:"name"`
	Parent string `bson:"parent" json:"parent"`
	Index  int    `bson:"index" json:"index"`
}

//文章结构
type Article_ struct {
	ArticleInfo
	ArticleContent
}

//文章信息
type ArticleInfo struct {
	Id    string `bson:"id"`
	Cat   string `bson: "cat"`
	Title string `bson: "title"`

	Subcat   string `bson: "subcat"`
	Subtitle string `bson: "subtitle"`
	ScanNum  int    `bson:"scannum"`
	LoveNum  int    `bson:"lovenum`
	CreateAt int64  `bson:"createdAt"`
	LastEdit int64  `bson:"lastedit"`
	Author   string `bson:"author"`
	Index    int    `bson:"index"`
}

//文章内容
type ArticleContent struct {
	Id      string `bson:"id"`
	Content string `bson: "content"`
}

//评论信息
type Comment struct {
	Id       string `bson:"id"`
	UserName string `bson:"username"`
	Time     int    `bson:"comtime"`
	Content  string `bson:"content"`
	LoveNum  int    `bson:"lovenum"`
	Parent   string `bson:"parent"`
	HeadIcon string `bson:"headicon"`
	ArtId    string `bson:"artid"`
}

//基础信息
type BaseInfo struct {
	Id   string `json:"id"  bson:"id"`
	Type string `json:"type" bson:"type"`
	Info string `json:"info" bson:"info"`
}

//访问信息
type VisitInfo struct {
	VisitNum int64 `json:"visitnum"`
}
