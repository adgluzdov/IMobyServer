package model

import (
	"io"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Uid string
	Scope int
}

type GetProfileInfoRequest struct {
	AuthenticationRequest
}

type GetProfileInfoResponse struct {
	Account
}

func NewProfileInfoRequest(reader io.Reader) (result *GetProfileInfoRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}