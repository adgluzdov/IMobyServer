package account

import (
	"github.com/adgluzdov/IMobyServer/model"
	"io"
	"encoding/json"
)

type GetProfileInfoRequest struct {
	model.AuthenticationRequest
}

type GetProfileInfoResponse struct {
	model.AccountInfo
}

func NewGetProfileInfoRequest(reader io.Reader) (result *GetProfileInfoRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}
