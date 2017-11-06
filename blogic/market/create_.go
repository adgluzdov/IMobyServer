package market

import (
	"github.com/adgluzdov/IMobyServer/model/market"
	"github.com/adgluzdov/IMobyServer/model"
	"github.com/adgluzdov/IMobyServer/data"
	"errors"
	"github.com/adgluzdov/IMobyServer/data/database"
	"github.com/adgluzdov/IMobyServer/blogic"
)

type Create_ struct {

}

func (Create_)Go(request *market.CreateRequest)(response market.CreateResponse,err error)  {
	db := database.GetMGOInstance()
	// Аутентификация
	var auth blogic.Authentication
	auth = new(blogic.Authentication_)
	authResponse,err := auth.Authenticate(request.AuthenticationRequest)
	if(err != nil){return }
	// Проверка Scope == MARKETER
	var account model.Account
	err = db.FindAccountById(authResponse.Id,&account)
	if(err != nil){return }
	if(account.Info.Scope != data.SCOPE_MARKETER) {
		err = errors.New("Account doesn't have scope")
		return
	}
	// Создание Market
	newMarket := model.Market{Id_Marketer:account.Id,Info:request.MarketInfo}
	newMarket.Id, err = db.CreateMarket(newMarket)
	if(err != nil){return }
	response.Id = newMarket.Id
	return
}
