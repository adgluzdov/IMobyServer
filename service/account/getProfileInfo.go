package account

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/model/account"
	blogic_account "github.com/adgluzdov/IMobyServer/blogic/account"
)

func GetProfileInfo(request *application.Request) (response *application.Response,err error) {
	requestBL, err := account.NewGetProfileInfoRequest(request.Data())
	if err!=nil {return application.NewResponse(err, nil)}

	var method blogic_account.GetProfileInfo
	method = new(blogic_account.GetProfileInfo_)
	responseBL, err := method.Go(requestBL)
	if err!=nil {return application.NewResponse(err.Error(),nil)}

	response,err = application.NewResponse(responseBL, err)
	return
}
