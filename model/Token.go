package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	AccountId bson.ObjectId
	Info      TokenInfo
}

type TokenInfo struct {
	Accsses_token string
	Expires_A int64
	Refresh_token string
	Expires_R int64
}

