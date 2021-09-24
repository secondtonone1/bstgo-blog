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
	Id    string `bson:"id" json:"infoid"`
	Cat   string `bson: "cat" json: "cat"`
	Title string `bson: "title" json: "title"`

	Subcat   string `bson: "subcat" json: "subcat"`
	Subtitle string `bson: "subtitle" json: "subtitle"`
	ScanNum  int    `bson:"scannum" json:"scannum"`
	LoveNum  int    `bson:"lovenum json:"lovenum`
	CreateAt int64  `bson:"createdAt" json:"createdAt"`
	LastEdit int64  `bson:"lastedit" json:"lastedit"`
	Author   string `bson:"author" json:"author"`
	Index    int    `bson:"index" json:"index"`
}

//文章内容
type ArticleContent struct {
	Id      string `bson:"id" json:"conid"`
	Content string `bson: "content" json: "content"`
}

//评论信息
type Comment struct {
	Id       string `bson:"id" json:"id"`
	UserName string `bson:"username" json:"username"`
	Time     int    `bson:"comtime" json:"comtime"`
	Content  string `bson:"content" json:"content"`
	LoveNum  int    `bson:"lovenum" json:"lovenum"`
	Parent   string `bson:"parent" json:"parent"`
	HeadIcon string `bson:"headicon" json:"headicon"`
	ArtId    string `bson:"artid" json:"artid"`
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

//二级菜单排序信息
type MenuSlice []*CatMenu

func (ms MenuSlice) Len() int {
	return len(ms)
}

func (ms MenuSlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms MenuSlice) Less(i, j int) bool {
	return ms[i].Index < ms[j].Index
}

//评论排序信息
type ComSlice []*Comment

func (cs ComSlice) Len() int {
	return len(cs)
}

func (cs ComSlice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

//从大到小
func (cs ComSlice) Less(i, j int) bool {
	return cs[i].Time > cs[j].Time
}

//文章首页排序
type HomeArtSlice []*Article_

func (hs HomeArtSlice) Len() int {
	return len(hs)
}

func (cs HomeArtSlice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

//从大到小
func (cs HomeArtSlice) Less(i, j int) bool {
	return cs[i].LastEdit > cs[j].LastEdit
}
