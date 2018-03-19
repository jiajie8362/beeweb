package main

import (
	"beeweb/common"
	"beeweb/routers"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	common.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}
	server.ListenAndServe()
}
