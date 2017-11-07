package market

import "github.com/adgluzdov/IMobyServer/model/market"

type GetById interface {
	Go(request *market.GetByIdRequest) (response market.GetByIdResponse, err error)
}
