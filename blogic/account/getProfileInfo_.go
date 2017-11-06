package account

import (
	"github.com/adgluzdov/IMobyServer/model/account"
	"github.com/adgluzdov/IMobyServer/data/database"
	"github.com/adgluzdov/IMobyServer/blogic"
)

type GetProfileInfo_ struct {

}

func (GetProfileInfo_)Go(request *account.GetProfileInfoRequest)(response account.GetProfileInfoResponse,err error)  {
	db := database.GetMGOInstance()
	// Аутентификация
	var auth blogic.Authentication
	auth = new(blogic.Authentication_)
	authResponse,err := auth.Authenticate(request.AuthenticationRequest)
	if(err != nil){return }
	// Поиск Аккаунта
	err = db.FindAccountById(authResponse.Id,&response)
	return
}
