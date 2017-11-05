package service

import (
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/blogic/auth"
	"github.com/adgluzdov/IMobyServer/application"
	blogic2 "github.com/adgluzdov/IMobyServer/blogic"
)

func Auth(request *application.Request) (response *application.Response,err error) {
	authRequest, err := model.NewAuthRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}
	var blogic blogic2.Auth
	blogic = new(blogic2.Auth_)
	authResponse, err := blogic.Authorize(authRequest)
	if err!=nil {return application.NewResponse(err.Error(),nil)}
	response,err = application.NewResponse(authResponse, err)
	return
}


