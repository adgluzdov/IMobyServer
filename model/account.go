package model

import (
	"io"
	"encoding/json"
)

type Account struct {
	Scope int
}

type ProfileInfoRequest struct {
	Accsses_token string
}

func NewProfileInfoRequest(reader io.Reader) (result *ProfileInfoRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}