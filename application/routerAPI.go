package application

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"webserver/filesystem"
	"webserver/http"
)

func getRouterAPI() Router {
	router := NewRouter("/api/memo")

	router.addHandler(GET, func(req http.HttpRequest, res http.HttpResponse) (http.HttpResponse, error) {
		fmt.Println("Get Memos")
		//Get all memo
		db := filesystem.GetDb()
		defer db.Close()

		memos, err := filesystem.GetAllMemo(db)

		if err != nil {
			return http.GetInternalErrorResponse(req, res), nil
		}
		//Create body
		body, err := json.Marshal(memos)
		if err != nil {
			return http.GetInternalErrorResponse(req, res), nil
		}
		//Return
		res.SetBody(body)

		res.StatusCode = "200"
		res.StatusMsg = "Success"

		fmt.Println(res.String())

		return res, nil
	})

	router.addHandler(DELETE, func(req http.HttpRequest, res http.HttpResponse) (http.HttpResponse, error) {
		//Get id from url
		idString := strings.TrimPrefix(req.URL, router.path+"/")
		idToDelete, err := strconv.Atoi(idString)
		if err != nil {
			return http.GetBadRequestResponse(req, res), nil
		}
		//Delete
		db := filesystem.GetDb()
		defer db.Close()

		err = filesystem.DeleteMemo(db, idToDelete)

		if err != nil {
			return http.GetInternalErrorResponse(req, res), nil
		}
		//Create body
		res.StatusCode = "200"
		res.StatusMsg = "Success"

		return res, nil
	})

	router.addHandler(POST, func(req http.HttpRequest, res http.HttpResponse) (http.HttpResponse, error) {
		memo, err := filesystem.ParseMemo(req.Body.Raw)

		if err != nil {
			fmt.Println(err.Error())
			return http.GetBadRequestResponse(req, res), nil
		}

		if memo.LastEditedBy != "" {
			//Do nothing
		} else if req.GetCookie("uuid").Name != "NOT_EXIST" {
			memo.LastEditedBy = req.GetCookie("uuid").Value
		} else {
			return http.GetBadRequestResponse(req, res), nil
		}

		db := filesystem.GetDb()
		defer db.Close()

		generatedMemo, err := filesystem.InsertMemo(db, memo)

		if err != nil {
			return http.GetInternalErrorResponse(req, res), nil
		}

		res.StatusCode = "200"
		res.StatusMsg = "Success"

		res.SetBody([]byte(generatedMemo.String()))

		fmt.Printf("Post: %s\n", generatedMemo.String())

		return res, nil
	})

	router.addHandler(PUT, func(req http.HttpRequest, res http.HttpResponse) (http.HttpResponse, error) {
		memo, err := filesystem.ParseMemo(req.Body.Raw)
		if err != nil {
			fmt.Println(err.Error())
			return http.GetBadRequestResponse(req, res), nil
		}
		if memo.LastEditedBy != "" {
			//Do nothing
		} else if req.GetCookie("uuid").Name != "NOT_EXIST" {
			memo.LastEditedBy = req.GetCookie("uuid").Value
		} else {
			fmt.Println("No uuid")
			return http.GetBadRequestResponse(req, res), nil
		}

		if memo.Id == 0 {
			return http.GetBadRequestResponse(req, res), nil
		}
		db := filesystem.GetDb()
		defer db.Close()

		err = filesystem.UpdateMemo(db, memo)
		if (err == filesystem.NoRowAffectedErr{}) {
			return http.GetNotFoundResponse(req), nil

		} else if err != nil {
			fmt.Println(err)

			return http.GetInternalErrorResponse(req, res), nil
		}

		res.StatusCode = "200"
		res.StatusMsg = "Success"

		return res, nil
	})

	return router
}
