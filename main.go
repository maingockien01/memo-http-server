package main

import "webserver/application"

func main() {

	webserver := application.NewWebserver("localhost", "8301")
	webserver.Start()
}
