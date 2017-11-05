package blogic

import (
	"github.com/adgluzdov/IMobyServer/model/market"
	"github.com/adgluzdov/IMobyServer/data/database"
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/data"
	"errors"
)

type Market interface {
	Create(request *market.CreateRequest)(response market.CreateResponse,err error)
	GetById(request *market.GetByIdRequest)(response market.GetByIdResponse,err error)
	Search(request *market.SearchRequest)(response market.SearchResponse,err error)
}

type Market_ struct {

}

func (Market_)Create(request *market.CreateRequest)(response market.CreateResponse,err error)  {
	var db database.MongoDB
	db.Init()
	defer db.Close()
	// Аутентификация
	var auth Authentication
	auth = new(Authentication_)
	authResponse,err := auth.Authenticate(request.AuthenticationRequest)
	if(err != nil){return }
	// Проверка Scope == MARKETER
	var account model.Account
	err = db.FindAccountById(authResponse.Id,&account)
	if(err != nil){return }
	if(account.Scope != data.SCOPE_MARKETER) {
		err = errors.New("Account doesn't have scope")
		return 
	}
	// Создание Market
	newMarket := market.Market{Id_Marketer:account.Id,MarketInfo:request.MarketInfo}
	newMarket.Id, err = db.CreateMarket(newMarket)
	if(err != nil){return }
	response.Id = newMarket.Id
	return
}

func (Market_)GetById(request *market.GetByIdRequest)(response market.GetByIdResponse,err error)  {
	return
}

func (Market_)Search(request *market.SearchRequest)(response market.SearchResponse,err error) {
	return
}