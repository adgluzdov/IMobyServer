package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
	"time"
	"errors"
	"github.com/adgluzdov/IMobyServer/data"
	"fmt"
	"github.com/adgluzdov/IMobyServer/data/database"
	"math/rand"
)

type Auth_ struct {

}

func (this *Auth_)Authorize(request *model.AuthorizationRequest)(response model.AuthorizationResponse,err error)  {
	//idToken -> tokenFB -> UID
	var auth AuthenticationFB
	auth = new(AuthenticationFB_)
	tokenFB, err := auth.Authenticate(model.AuthenticationFBRequest{request.IdToken})
	if err != nil {return }
	if(time.Now().Unix() > tokenFB.Expires) {
		err = errors.New("Expired token")
		return
	}
	UID := tokenFB.UID
	// есть ли Account с данным uid
	db := database.GetMongoDB()
	var account model.Account
	findError := db.FindAccount(UID,&account)
	if(findError != nil){
		// create User
		account.Uid = tokenFB.UID
		account.Info.Scope = data.SCOPE_USER
		db.InsertAccount(account)
		return
	}else {
		// create token
		var token model.Token
		token.AccountId = account.Id
		token.Info.Accsses_token = randToken()
		token.Info.Refresh_token = randToken()
		token.Info.Expires_A = time.Now().Unix() + data.Expires_A
		token.Info.Expires_R = time.Now().Unix() + data.Expires_R
		// insert token into DB
		err = db.InsertToken(token)
		if err!=nil {return }
		// token -> AuthorizationResponse
		response.Token = token.Info
		return
	}
	return
}

func randToken() string {
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
