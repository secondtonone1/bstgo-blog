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
	CatId        string        `bson:"catid"`
	Name         string        `bson:"name"`
	SubCatMenus_ []*SubCatMenu `bson:"subcatmenus"`
}

//子分类目录
type SubCatMenu struct {
	SubCatId string `bson:"subcatid"`
	Name     string `bson:"name"`
}

//总目录
type Menu struct {
	CatMenus_ []*CatMenu `bson:"catmenus"`
}
