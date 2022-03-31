package application

import (
	"webserver/http"

	id "github.com/google/uuid"
)

func getMiddlewareTrackingId() func(http.HttpRequest, http.HttpResponse) (http.HttpResponse, error) {
	return func(req http.HttpRequest, res http.HttpResponse) (http.HttpResponse, error) {
		//TODO: make uuid random?
		uuid := req.GetCookie("uuid")
		if uuid.Name == "NOT_EXIST" {
			uuid = http.Cookie{
				Name:  "uuid",
				Value: id.NewString(),
				Path:  "/",
			}
			res.AddCookie(uuid)
			// req.AddCookie(uuid)
		}

		return res, nil
	}
}
