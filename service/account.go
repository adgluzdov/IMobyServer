package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/blogic"
)

func GetProfileInfo(request *application.Request) (response *application.Response,err error) {
	profileInfoRequest, err := model.NewProfileInfoRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}
	var account blogic.Account
	account = new(blogic.Account_)
	profileInfoResponse, err := account.GetProfileInfo(profileInfoRequest)
	if err!=nil {return application.NewResponse(err.Error(),nil)}
	response,err = application.NewResponse(profileInfoResponse, err)
	return
}
