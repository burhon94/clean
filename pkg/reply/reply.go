package reply

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/burhon94/clean/internal/structs"
)

func ReplyJSON(w http.ResponseWriter, resp *structs.Response) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(http.StatusOK) //service status always be OK

	_, err = w.Write(data)
	if err != nil {
		log.Printf("error when send reply: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
