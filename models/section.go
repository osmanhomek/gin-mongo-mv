package models

import "gopkg.in/mgo.v2/bson"

type (
	// Section represents the structure of our resource
	Section struct {
		Id         bson.ObjectId `bson:"_id"`
		orderIndex int           `bson:"orderindex"`
		engName    string        `bson:"engname"`
		trName     string        `bson:"trname"`
		avatar     string        `bson:"avatar"`
	}
)
