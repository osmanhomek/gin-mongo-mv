package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	Section struct {
		Id      bson.ObjectId `json:"id" bson:"_id"`
		EngName string        `json:"eng_name" bson:"eng_name"`
		TrName  string        `json:"tr_name" bson:"tr_name"`
		Avatar  string        `json:"avatar" bson:"avatar"`
	}
)
