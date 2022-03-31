package application

import (
	"os"
	"webserver/http"
)

// Default Router
// Get file from file system
func getRouterDefault() (router Router) {
	router = NewRouter("")
	router.addHandler(GET, func(req http.HttpRequest, res http.HttpResponse) (http.HttpResponse, error) {
		path := "." + req.URL

		if req.URL == "/" {
			path = "./index.html"
		}

		file, err := os.ReadFile(path)
		if err != nil {
			return http.GetNotFoundResponse(req), nil
		}

		res.StatusCode = "200"
		res.StatusMsg = "Success"
		res.SetBody(file)

		return res, nil
	})

	return

}
