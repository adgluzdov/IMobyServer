package model

import (
	"io"
	"encoding/json"
)

type AuthRequest struct {
	IdToken string
}

type AuthResponse struct {
	Token TokenIM
}

type TokenIM_DB struct {
	Uid   string
	Token TokenIM
}

type TokenIM struct {
	Accsses_token string
	Expires_A int64
	Refresh_token string
	Expires_R int64
}

func NewAuthRequest(reader io.Reader) (result *AuthRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}


