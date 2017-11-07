package market

import "github.com/adgluzdov/IMobyServer/model/market"

type Search interface {
	Go(request *market.SearchRequest) (response market.SearchResponse, err error)
}