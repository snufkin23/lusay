package http

import (
	"encoding/json"
	"net/http"
)

type errorBody struct {
	Error string `json:"error"`
}

func respond(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func respondErr(w http.ResponseWriter, status int, msg string) {
	respond(w, status, errorBody{Error: msg})
}
