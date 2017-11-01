package auth

import (
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
)

type Firebase struct {
	client *auth.Client
}

func (this *Firebase)Init()(err error){
	opt := option.WithCredentialsFile("imoby-d96c5-firebase-adminsdk-chv6j-c88bfa2694.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	this.client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return
}

func (this *Firebase)VerifyIDToken(idToken string)(token *auth.Token,err error)  {
	token, err = this.client.VerifyIDToken(idToken)
	return
}

func (this *Firebase)CustomToken(uid string)(token string,err error)  {
	token, err = this.client.CustomToken(uid)
	return
}