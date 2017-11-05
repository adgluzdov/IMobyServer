package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/data/database"
	"time"
	"errors"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"log"
	"context"
)

type Authentication interface {
	Authenticate(request model.AuthenticationRequest)(response model.AuthenticationResponse, err error)
}

type Authentication_ struct {

}

func (this *Authentication_)Authenticate(request model.AuthenticationRequest)(response model.AuthenticationResponse, err error)  {
	var db database.MongoDB
	db.Init()
	defer db.Close()
	var token model.TokenIM_DB
	err = db.FindToken(request.Accsses_token,&token)
	if(err != nil) {return }
	if(token.Token.Expires_A < time.Now().Unix()) {
		err = errors.New("Token has expired.")
	}
	// token -> response
	response.Id = token.UserId
	return
}

type AuthenticationFB interface {
	Authenticate(request model.AuthenticationFBRequest)(response model.AuthenticationFBResponse, err error)
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