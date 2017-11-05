package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Market struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Id_Marketer bson.ObjectId
	Info MarketInfo
}
type MarketInfo struct {
	Name string
	Description string
}
