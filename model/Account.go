package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Uid string
	Info AccountInfo
}

type AccountInfo struct {
	Scope int
}

