package main

import (
	"encoding/json"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "pong",
	}

	json.NewEncoder(w).Encode(response)
}
