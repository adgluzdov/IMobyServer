package model

import (
	"io"
	"encoding/json"
)

type Model struct {
	Hello string
}

func NewModel(reader io.Reader) (result *Model, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}