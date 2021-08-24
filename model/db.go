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
type Article struct {
	Id       string `bson:"id"`
	Cat      string `bson: "cat"`
	Title    string `bson: "title"`
	Content  string `bson: "content"`
	Subcat   string `bson: "subcat"`
	Subtitle string `bson: "subtitle"`
	ScanNum  int    `bson:"scannum"`
	LoveNum  int    `bson:"lovenum`
	CreateAt int64  `bson:"createdAt"`
	LastEdit int64  `bson:"lastedit"`
	Author   string `bson:"author"`
	Index    int    `bson:"index"`
}
