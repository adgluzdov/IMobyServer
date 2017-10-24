package model

import (
	"io"
	"encoding/json"
)

type TokenId struct {
	Id string
}

func NewTokenId(reader io.Reader) (result *TokenId, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}
