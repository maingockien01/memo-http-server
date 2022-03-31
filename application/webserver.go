package application

import (
	"fmt"
	"strings"
	"webserver/http"
)

type Webserver struct {
	httpServer http.HttpServer
	Host       string
	Port       string
	routers    []Router
}

func NewWebserver(host string, port string) Webserver {
	webserver := Webserver{
		Host: host,
		Port: port,
		httpServer: http.HttpServer{
			Host: host,
			Port: port,
		},
	}

	webserver.httpServer.Handler = webserver.RunApplication

	webserver.httpServer.Setup()

	//Add routers
	webserver.routers = append(webserver.routers, getRouterAPI())
	//Defaul router is last one - always
	webserver.routers = append(webserver.routers, getRouterDefault())

	return webserver
}

func (server *Webserver) Start() {
	fmt.Println("Starting server ...")
	server.httpServer.Start()
	for {
		//Do nothing just keep the thread alive}
	}
}

func (webserver *Webserver) RunApplication(req http.HttpRequest, res http.HttpResponse) (http.HttpResponse, error) {
	var router Router

	trackingId := getMiddlewareTrackingId()
	res, _ = trackingId(req, res)

	for _, avaialbeRouter := range webserver.routers {
		if strings.HasPrefix(req.URL, avaialbeRouter.path) {
			router = avaialbeRouter
			break
		}
	}

	return router.handler(req, res)
}
