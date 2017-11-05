package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
	"time"
	"github.com/adgluzdov/IMobyServer/data"
	"fmt"
	"errors"
	"math/rand"
	"github.com/adgluzdov/IMobyServer/data/database"
)

type Auth interface {
	Authorize(request *model.AuthRequest)(response model.AuthResponse,err error)
}

type Auth_ struct {}

func (this *Auth_)Authorize(request *model.AuthRequest)(response model.AuthResponse,err error)  {
	//idToken -> tokenFB
	var auth AuthenticationFB
	auth = new(AuthenticationFB_)
	tokenFB, err := auth.Authenticate(model.AutenticationFBRequest{request.IdToken})
	if err != nil {return }
	if(time.Now().Unix() > tokenFB.Expires) {
		err = errors.New("Expired token")
		return
	}
	UID := tokenFB.UID

	var db database.MongoDB
	db.Init()
	defer db.Close()
	var account model.Account
	findError := db.FindAccount(UID,&account)
	if(findError != nil){
		// create User
		account.Uid = tokenFB.UID
		account.Scope = data.SCOPE_USER
		db.InsertAccount(account)
		return
	}else {
		// create tokenIM_DB
		var tokenIM_DB model.TokenIM_DB
		tokenIM_DB.Uid = tokenFB.UID
		tokenIM_DB.Token.Accsses_token = randToken()
		tokenIM_DB.Token.Refresh_token = randToken()
		tokenIM_DB.Token.Expires_A = time.Now().Unix() + data.Expires_A
		tokenIM_DB.Token.Expires_R = time.Now().Unix() + data.Expires_R
		// insert tokenIM_DB into DB
		err = db.InsertTokenIM(tokenIM_DB)
		if err!=nil {return }
		// tokenIM_DB -> AuthResponse
		response.Token = tokenIM_DB.Token
		return
	}
	return
}

func randToken() string {
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}