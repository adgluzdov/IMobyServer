package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
	"time"
	"errors"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"log"
	"github.com/adgluzdov/IMobyServer/data/database"
	"context"
)

type Authentication_ struct {

}

func (this *Authentication_)Authenticate(request model.AuthenticationRequest)(response model.AuthenticationResponse, err error)  {
	db := database.GetMongoDB()
	var token model.Token
	err = db.FindToken(request.Accsses_token,&token)
	if(err != nil) {return }
	if(token.Info.Expires_A < time.Now().Unix()) {
		err = errors.New("TokenInfo has expired.")
	}
	// token -> response
	response.Id = token.AccountId
	return
}

type AuthenticationFB_ struct {
	client *auth.Client
}

func (this *AuthenticationFB_)Authenticate(request model.AuthenticationFBRequest)(response model.AuthenticationFBResponse, err error){
	opt := option.WithCredentialsFile("imoby-d96c5-firebase-adminsdk-chv6j-c88bfa2694.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	this.client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	response.Token, err = this.client.VerifyIDToken(request.TokenId)
	return
}
