package market

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model/market"
	blogic_market "github.com/adgluzdov/IMobyServer/blogic/market"
)

func GetById(request *application.Request) (response *application.Response,err error) {
	requestBL, err := market.NewGetByIdRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}

	var method blogic_market.GetById
	method = new(blogic_market.GetById_)
	responseBL, err := method.Go(requestBL)
	if err!=nil {return application.NewResponse(err.Error(),nil)}

	response,err = application.NewResponse(responseBL, err)
	return
}
