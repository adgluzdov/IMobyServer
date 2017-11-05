package market

import (
	"github.com/adgluzdov/IMobyServer/model"
	"gopkg.in/mgo.v2/bson"
	"io"
	"encoding/json"
)

type CreateRequest struct {
	model.AuthenticationRequest
	model.MarketInfo
}

type CreateResponse struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

func NewCreateRequest(reader io.Reader) (result *CreateRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}