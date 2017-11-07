package market

import (
	"github.com/adgluzdov/IMobyServer/model"
	"io"
	"encoding/json"
)

type SearchRequest struct{
	Query string
	Sort int
	model.AuthenticationRequest
}

type SearchResponse struct{
	Markets []model.Market
}

func NewSearchRequest(reader io.Reader) (result *SearchRequest, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}
