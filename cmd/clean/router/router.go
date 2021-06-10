package router

import (
	"net/http"

	"github.com/burhon94/clean/cmd/clean/handlers"
)

func InitRouter(prefix string) *http.ServeMux {

	newServeMux := http.NewServeMux()

	newServeMux.HandleFunc(prefix+"/ping", handlers.HandlerPing)

	return newServeMux
}
