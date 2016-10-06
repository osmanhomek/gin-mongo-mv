package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	Section struct {
		Id      bson.ObjectId `json:"id" bson:"_id"`
		orderIndex int        `json:"order_index" bson:"order_index"`
		engName string        `json:"eng_name" bson:"eng_name"`
		trName  string        `json:"tr_name" bson:"tr_name"`
		avatar  string        `json:"avatar" bson:"avatar"`
	}
)
