package market

import (
	"github.com/adgluzdov/IMobyServer/model/market"
	"github.com/adgluzdov/IMobyServer/blogic"
	"github.com/adgluzdov/IMobyServer/data/database"
	"github.com/adgluzdov/IMobyServer/model"
)

type GetById_ struct {

}

func (GetById_)Go(request *market.GetByIdRequest) (response market.GetByIdResponse, err error) {
	db := database.GetMongoDB()
	// Аутентификация
	var auth blogic.Authentication
	auth = new(blogic.Authentication_)
	_,err = auth.Authenticate(request.AuthenticationRequest)
	if(err != nil){return }
	// Получаем группу ипо id
	var market model.Market
	err = db.FindMarketById(request.Id_Market,&market)
	if(err != nil){return }
	response.Market = market
	return
}