package market

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/adgluzdov/IMobyServer/model"
	"io"
	"encoding/json"
)

type GetByIdRequest struct{
	Id_Market bson.ObjectId
	model.AuthenticationRequest
}

type GetByIdResponse struct{
	model.Market
}

func NewGetByIdRequest(reader io.Reader) (result *GetByIdRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}
