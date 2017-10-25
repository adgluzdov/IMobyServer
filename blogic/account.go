package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/data/database"
	"github.com/adgluzdov/IMobyServer/auth"
)

func GetProfileInfo(request *model.ProfileInfoRequest)(response model.Account,err error)  {
	var db database.MongoDB
	db.Init()
	defer db.Close()
	// Аутентификация
	var authRequest model.IMobyAuthRequest
	authRequest.Init(request.Accsses_token)
	authResponse,err := auth.AuthenticationIM(db,authRequest)
	if(err != nil){return }
	// Поиск Аккаунта
	err = db.FindAccount(authResponse.Uid,&response)
	return
}