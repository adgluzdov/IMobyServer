package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/data/database"
)

func GetProfileInfo(request *model.ProfileInfoRequest)(response model.Account,err error)  {
	var db database.MongoDB
	db.Init()
	defer db.Close()
	// Аутентификация
	var auth Authentication
	auth = new(Authentication_)
	authRequest := model.AutenticationRequest{request.Accsses_token}
	authResponse := model.AutenticationResponse{}
	authResponse,err = auth.Authenticate(authRequest)
	if(err != nil){return }
	// Поиск Аккаунта
	err = db.FindAccount(authResponse.Uid,&response)
	return
}