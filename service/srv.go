package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model"
)

func Service(request *application.Request) (*application.Response, error) {
	model, err := model.NewModel(request.Data())
	if err!=nil {return application.NewResponse(nil, err)}
	result := model
	return application.NewResponse(result, err)
}
