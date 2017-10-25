package model

import (
	"io"
	"encoding/json"
)

type AuthLP struct {
	Login string
	Password string
}

type AuthResponse struct {
	AccessToken string
}

func NewAuthLP(reader io.Reader) (result *AuthLP, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}
