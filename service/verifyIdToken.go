package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/auth"
	"gopkg.in/mgo.v2/bson"
)

func VerifyIdToken(request *application.Request) (response *application.Response,err error) {
	tokenId, err := model.NewTokenId(request.Data())
	if err!=nil {return application.NewResponse(nil, err)}
	var auth auth.Firebase
	err = auth.Init()
	if err!=nil {return application.NewResponse(nil, err)}
	token,err := auth.VerifyIDToken(tokenId.Id)
	if err!=nil {return application.NewResponse(nil, err)}
	if err!=nil {
		response,err = application.NewResponse(bson.M{"result":"invalid"}, err)
	} else {
		response,err = application.NewResponse(bson.M{"result":token}, err)
	}
	return
}

