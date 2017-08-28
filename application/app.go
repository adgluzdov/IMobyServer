package application

import (
	"log"
	"net/http"
	"fmt"
)

type Routes map[string]func(*Request) (*Response, error)

type application struct {
	routes Routes
	host string
	port string
}

func Init(routes Routes) *application {
	result := new(application)
	result.routes = routes
	result.host = ""
	result.port = "8001"
	return result
}

func (app *application) Run() {
	for url, method := range app.routes {
		http.HandleFunc(url, httpserve(method))
	}
	listen := fmt.Sprintf("%s:%s", app.host, app.port)
	log.Println("Server started.")
	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Println(err)
	}
}

func httpserve(handler func(*Request) (*Response, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		response, err := handler(NewRequest(req))
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		w.Header().Set(response.HeaderKey, response.HeaderValue);
		w.Write(response.Result)
	}
}
