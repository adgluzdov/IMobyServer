package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/blogic"
	"github.com/adgluzdov/IMobyServer/model/market"
)

func Create(request *application.Request) (response *application.Response,err error) {
	createRequest, err := market.NewCreateRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}
	var market blogic.Market
	market = new(blogic.Market_)
	createResponse, err := market.Create(createRequest)
	if err!=nil {return application.NewResponse(err.Error(),nil)}
	response,err = application.NewResponse(createResponse, err)
	return
}

