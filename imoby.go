package main

import (
	"github.com/adgluzdov/IMobyServer/application"
	"github.com/adgluzdov/IMobyServer/service"
)

func main()  {
	Application := application.Init(application.Routes{"/calculate": service.Service})
	Application.Run()
}
