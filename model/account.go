package model

import (
	"io"
	"encoding/json"
)

type Account struct {
	Uid string
	Scope int
}

type GetProfileInfoRequest struct {
	Accsses_token string
}

type GetProfileInfoResponse struct {
	Account
}

func NewProfileInfoRequest(reader io.Reader) (result *GetProfileInfoRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}