package blogic

import (
	"github.com/adgluzdov/IMobyServer/model"
)

type Auth interface {
	Authorize(request *model.AuthorizationRequest)(response model.AuthorizationResponse,err error)
}