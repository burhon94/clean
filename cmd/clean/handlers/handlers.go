package handlers

import (
	"github.com/burhon94/clean/pkg/reply"
	"net/http"

	"github.com/burhon94/clean/internal/structs"
)

//Here only http-Req-Resp-logic
func HandlerPing(w http.ResponseWriter, r *http.Request) {
	//if have business logic realised it to /clean/internal/bizLogic

	var resp = structs.Response{
		Code:    202,
		Payload: "pong",
		Message: "Success",
	}

	defer reply.ReplyJSON(w, &resp)
}
