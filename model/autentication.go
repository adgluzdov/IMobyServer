package model

import "firebase.google.com/go/auth"

type AutenticationRequest struct {
	Accsses_token string
}

type AutenticationResponse struct {
	Uid string
}

type AutenticationFBRequest struct {
	TokenId string
}

type AutenticationFBResponse struct {
	*auth.Token
}