package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/data/database"
)

type Account interface {
	GetProfileInfo(request *model.GetProfileInfoRequest)(response model.Account,err error)
}

type Account_ struct {

}

func (Account_)GetProfileInfo(request *model.GetProfileInfoRequest)(response model.Account,err error)  {
	db := database.GetMGOInstance()
	defer db.Close()
	// Аутентификация
	var auth Authentication
	auth = new(Authentication_)
	authResponse,err := auth.Authenticate(request.AuthenticationRequest)
	if(err != nil){return }
	// Поиск Аккаунта
	err = db.FindAccountById(authResponse.Id,&response)
	return
}