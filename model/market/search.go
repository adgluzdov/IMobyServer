package market

import "github.com/adgluzdov/IMobyServer/model"

type SearchRequest struct{
	Query string
	Sort int
	model.AuthenticationRequest
}

type SearchResponse struct{
	Markets []model.Market
}
