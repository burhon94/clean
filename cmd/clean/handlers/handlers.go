package handlers

import (
	"encoding/json"
	"net/http"
)

//TODO move all structs to /clean/internal/structs/...
type responseStruct struct {
	Code    int         `json:"code"`
	Payload interface{} `json:"payload"`
	Message string      `json:"message"`
}

//Here only http-Req-Resp-logic
func HandlerPing(w http.ResponseWriter, r *http.Request) {
	//if have business logic realised it to /clean/internal/bizLogic

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
