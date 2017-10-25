package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/auth"
	"github.com/adgluzdov/IMobyServer/data/database"
	"fmt"
	"math/rand"
	"github.com/adgluzdov/IMobyServer/data"
	"time"
)

func Auth(request *model.AuthRequest)(response model.AuthResponse,err error)  {
	//idToken -> token
	var auth auth.Firebase
	err = auth.Init()
	if err != nil {return }
	token, err := auth.VerifyIDToken(request.IdToken)
	if err != nil {return }

	var db database.MongoDB
	db.Init()
	defer db.Close()
	var account model.Account
	findError := db.FindAccount(token.UID,&account)
	if(findError != nil){
		account.Scopes = append(account.Scopes, data.SCOPE_USER)
		db.InsertAccount(account)
		return
	}else {
		// create tokenIMoby
		var tokenIMoby model.TokenIMoby
		tokenIMoby.Uid = token.UID
		tokenIMoby.Token.Accsses_token = randToken()
		tokenIMoby.Token.Refresh_token = randToken()
		tokenIMoby.Token.Expires_A = time.Now().Unix() + data.Expires_A
		tokenIMoby.Token.Expires_R = time.Now().Unix() + data.Expires_R
		// insert tokenIMoby into DB
		err = db.InsertTokenIMoby(tokenIMoby)
		if err!=nil {return }
		// tokenIMoby -> AuthResponse
		response.Token = tokenIMoby.Token
		return
	}
	return
}

func randToken() string {
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}