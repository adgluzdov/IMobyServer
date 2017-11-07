package market

import (
	"github.com/adgluzdov/IMobyServer/model/market"
	"github.com/adgluzdov/IMobyServer/blogic"
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/data/database"
)

type Search_ struct {

}

func (Search_)Go(request *market.SearchRequest) (response market.SearchResponse, err error) {
	db := database.GetMongoDB()
	// Аутентификация
	var auth blogic.Authentication
	auth = new(blogic.Authentication_)
	_,err = auth.Authenticate(request.AuthenticationRequest)
	if(err != nil){return }
	// Получаем группу ипо id
	var markets []model.Market
	err = db.FindMarketAll(&markets)
	if(err != nil){return }
	response.Markets = markets
	return
}
