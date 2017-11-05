package model

import (
	"firebase.google.com/go/auth"
	"gopkg.in/mgo.v2/bson"
)

type AuthenticationRequest struct {
	Accsses_token string
}

type AuthenticationResponse struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

type AuthenticationFBRequest struct {
	TokenId string
}

type AuthenticationFBResponse struct {
	*auth.Token
}