package router

import (
	"github.com/burhon94/clean/cmd/clean/handlers"
	"net/http"
)

func InitRouter(prefix string) *http.ServeMux {

	newServeMux := http.NewServeMux()

	newServeMux.HandleFunc(prefix+"/ping", handlers.HandlerPing)

	return newServeMux
}


