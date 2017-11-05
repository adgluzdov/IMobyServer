package market

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/adgluzdov/IMobyServer/model"
)

type GetByIdRequest struct{
	Id_Market bson.ObjectId
	model.AuthenticationRequest
}

type GetByIdResponse struct{
	Market
}
