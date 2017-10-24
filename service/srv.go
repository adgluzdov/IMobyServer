package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/data"
	"log"
)

func Service(request *application.Request) (*application.Response, error) {
	log.Print("Tc")
	model, err := model.NewModel(request.Data())
	if err!=nil {return application.NewResponse(nil, err)}
	session, err := data.InitSession();
	collection := session.DB("test").C("test")
	err = collection.Insert(model);
	if model.Hello != "Hello" {model.Hello = "Поприветствуй меня"}
	return application.NewResponse(model, err)
}
