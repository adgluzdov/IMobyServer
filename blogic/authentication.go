package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
)

type Authentication interface {
	Authenticate(request model.AuthenticationRequest)(response model.AuthenticationResponse, err error)
}

type AuthenticationFB interface {
	Authenticate(request model.AuthenticationFBRequest)(response model.AuthenticationFBResponse, err error)
}