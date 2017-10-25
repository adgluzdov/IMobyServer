package main

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/service"
)

func main()  {
	Application := application.Init(application.Routes{"/srv": service.Service,"/auth": service.Auth})
	Application.Run()
}
