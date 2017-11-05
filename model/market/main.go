package market

import (
	"gopkg.in/mgo.v2/bson"
)

type Market struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Id_Marketer bson.ObjectId
	MarketInfo
}
type MarketInfo struct {
	Name string
	Description string
}
