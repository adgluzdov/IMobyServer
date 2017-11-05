package service

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/blogic"
	"github.com/adgluzdov/IMobyServer/model/account"
)

func GetProfileInfo(request *application.Request) (response *application.Response,err error) {
	getProfileInfoRequest, err := account.NewGetProfileInfoRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}
	var account2 blogic.Account
	account2 = new(blogic.Account_)
	profileInfoResponse, err := account2.GetProfileInfo(getProfileInfoRequest)
	if err!=nil {return application.NewResponse(err.Error(),nil)}
	response,err = application.NewResponse(profileInfoResponse, err)
	return
}
