package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/blogic"
)

func Auth(request *application.Request) (response *application.Response,err error) {
	authRequest, err := model.NewAuthRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}
	authResponse, err := blogic.Auth(authRequest)
	if err!=nil {return application.NewResponse(err.Error(),nil)}
	response,err = application.NewResponse(authResponse, err)
	return
}


