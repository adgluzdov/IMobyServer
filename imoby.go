package main

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/service"
	"github.com/adgluzdov/IMobyServer/service/account"
	"github.com/adgluzdov/IMobyServer/service/market"
)

func main()  {
	Application := application.Init(application.Routes{"/auth": service.Auth,"/account.getProfileInfo": account.GetProfileInfo,"/market.create": market.Create})
	Application.Run()
}
