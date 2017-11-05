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
	var db database.MongoDB
	db.Init()
	defer db.Close()
	// Аутентификация
	var auth Authentication
	auth = new(Authentication_)
	authRequest := model.AutenticationRequest{request.Accsses_token}
	authResponse,err := auth.Authenticate(authRequest)
	if(err != nil){return }
	// Поиск Аккаунта
	err = db.FindAccount(authResponse.Uid,&response)
	return
}