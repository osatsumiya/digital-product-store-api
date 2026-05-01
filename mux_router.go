package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MuxHello(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello from Gorilla Mux",
	}
	json.NewEncoder(w).Encode(response)
}

func SetupMuxRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/mux/hello", MuxHello).Methods("GET")

	return r
}
