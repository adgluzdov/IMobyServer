package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/blogic"
)

func Auth(request *application.Request) (response *application.Response,err error) {
	authRequest, err := model.NewAuthLP(request.Data())
	if err!=nil {return application.NewResponse(nil, err)}
	authResponse, err := blogic.Auth(authRequest)
	//
	//var auth auth.Firebase
	//err = auth.Init()
	//if err!=nil {return application.NewResponse(nil, err)}
	//token,err := auth.VerifyIDToken(tokenId.Id)
	//if err!=nil {return application.NewResponse(nil, err)}
	//if err!=nil {
	//	response,err = application.NewResponse(bson.M{"result":"invalid"}, err)
	//} else {
	//	response,err = application.NewResponse(bson.M{"result":token}, err)
	//}
	response,err = application.NewResponse(authResponse, err)
	return
}


