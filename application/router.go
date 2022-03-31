package application

import "webserver/http"

type Router struct {
	path     string
	handlers map[string]func(http.HttpRequest, http.HttpResponse) (http.HttpResponse, error)
}

const GET = "GET"
const POST = "POST"
const DELETE = "DELETE"
const PUT = "PUT"

func NewRouter(path string) Router {
	router := Router{
		path:     path,
		handlers: make(map[string]func(http.HttpRequest, http.HttpResponse) (http.HttpResponse, error)),
	}

	return router
}

func (router *Router) addHandler(method string, handler func(http.HttpRequest, http.HttpResponse) (http.HttpResponse, error)) {
	router.handlers[method] = handler
}

func (router *Router) getHandler(method string) func(http.HttpRequest, http.HttpResponse) (http.HttpResponse, error) {
	return router.handlers[method]
}

func (router *Router) handler(req http.HttpRequest, res http.HttpResponse) (http.HttpResponse, error) {
	method := req.Method

	handler := router.getHandler(method)

	if handler == nil {
		return http.GetNotFoundResponse(req), nil
	}

	return handler(req, res)
}
