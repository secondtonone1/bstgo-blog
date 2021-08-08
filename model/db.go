package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Session struct {
	Id  primitive.ObjectID `json:"id" bson:"_id"`
	Sid string             `json:"sid" bson:"sid"`
}
