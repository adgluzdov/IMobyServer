package market

import (
	"github.com/adgluzdov/IMobyServer/application"
	blogic_market "github.com/adgluzdov/IMobyServer/blogic/market"
	"github.com/adgluzdov/IMobyServer/model/market"
)

func Create(request *application.Request) (response *application.Response,err error) {
	requestBL, err := market.NewCreateRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}

	var method blogic_market.Create
	method = new(blogic_market.Create_)
	responseBL, err := method.Go(requestBL)
	if err!=nil {return application.NewResponse(err.Error(),nil)}

	response,err = application.NewResponse(responseBL, err)
	return
}

