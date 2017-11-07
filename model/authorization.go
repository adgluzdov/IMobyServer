package model

import (
	"io"
	"encoding/json"
)

type AuthorizationRequest struct {
	IdToken string
}

type AuthorizationResponse struct {
	Token TokenInfo
}

func NewAuthorizationRequest(reader io.Reader) (result *AuthorizationRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}

