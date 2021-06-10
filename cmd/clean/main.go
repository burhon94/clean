package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	var (
		prefix = "/api"
		addr   = ":9999"
	)

	newMux := initRoute(prefix)

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

func initRoute(prefix string) *http.ServeMux {

	newServeMux := http.NewServeMux()

	newServeMux.HandleFunc(prefix+"/ping", pingHandler)

	return newServeMux
}

type responseStruct struct {
	Code    int         `json:"code"`
	Payload interface{} `json:"payload"`
	Message string      `json:"message"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	var resp = responseStruct{
		Code:    202,
		Payload: "pong",
		Message: "Success",
	}

	w.Header().Set("Content-Type", "application/json")


	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}
