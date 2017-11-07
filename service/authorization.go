package service

import (
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/blogic"
)

func Auth(request *application.Request) (response *application.Response,err error) {
	requestBL, err := model.NewAuthorizationRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}

	var method blogic.Auth
	method = new(blogic.Auth_)
	responseBL, err := method.Authorize(requestBL)
	if err!=nil {return application.NewResponse(err.Error(),nil)}

	response,err = application.NewResponse(responseBL, err)
	return
}


