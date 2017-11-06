package market

import "github.com/adgluzdov/IMobyServer/model/market"

type Create interface {
	Go(request *market.CreateRequest) (response market.CreateResponse, err error)
}
