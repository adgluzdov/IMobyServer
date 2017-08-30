package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model"
	"log"
)

func Service(request *application.Request) (*application.Response, error) {
	model, err := model.NewModel(request.Data())
	log.Print(model)
	if err!=nil {return application.NewResponse(nil, err)}
	if model.Hello != "Hello" {model.Hello = "Поприветствуй меня"}
	return application.NewResponse(model, err)
}
