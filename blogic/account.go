package blogic

import (
	"github.com/adgluzdov/IMobyServer/model/account"
	"github.com/adgluzdov/IMobyServer/data/database"
)

type Account interface {
	GetProfileInfo(request *account.GetProfileInfoRequest)(response account.GetProfileInfoResponse,err error)
}

type Account_ struct {

}

func (Account_)GetProfileInfo(request *account.GetProfileInfoRequest)(response account.GetProfileInfoResponse,err error)  {
	db := database.GetMGOInstance()
	// Аутентификация
	var auth Authentication
	auth = new(Authentication_)
	authResponse,err := auth.Authenticate(request.AuthenticationRequest)
	if(err != nil){return }
	// Поиск Аккаунта
	err = db.FindAccountById(authResponse.Id,&response)
	return
}