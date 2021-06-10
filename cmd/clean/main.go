package main

import (
	"log"
	"net/http"

	"github.com/burhon94/clean/cmd/clean/router"
)

func main() {
	var (
		prefix = "/api"
		addr   = ":9999"
	)

	newMux := router.InitRouter(prefix)

	newServer := http.Server{
		Handler: newMux,
		Addr:    addr,
	}

	defer func() {

		if err2 := recover(); err2 != nil {
			log.Println(err2)
			return
		}

		newServer.ListenAndServe()

	}()
}
