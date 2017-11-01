package auth

import (
	"github.com/adgluzdov/IMobyServer/data/database"
	"github.com/adgluzdov/IMobyServer/model"
)


func AuthenticationIM(db database.MongoDB,request model.IMobyAuthRequest)(response model.IMobyAuthResponse,err error)  {
	err = db.FindUid(request.Accsses_token,&response)
	return
}