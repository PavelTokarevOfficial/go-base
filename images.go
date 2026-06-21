package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func handleImages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleCreateImage(w)
	case http.MethodGet:
		handleListImages(w)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleImageByID(w http.ResponseWriter, r *http.Request) {
	id, err := imageIDFromRequest(r)
	if err != nil {
		http.Error(w, "invalid image id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGetImage(w, id)
	case http.MethodPost:
		handleUpdateImage(w, id)
	case http.MethodDelete:
		handleDeleteImage(w, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func imageIDFromRequest(r *http.Request) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/images/")
	return strconv.Atoi(idStr)
}

func handleCreateImage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "create image",
	}

	json.NewEncoder(w).Encode(response)
}

func handleListImages(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "get all images",
	}

	json.NewEncoder(w).Encode(response)
}

func handleUpdateImage(w http.ResponseWriter, id int) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "update image " + strconv.Itoa(id),
	}

	json.NewEncoder(w).Encode(response)
}

func handleGetImage(w http.ResponseWriter, id int) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "get image by id " + strconv.Itoa(id),
	}

	json.NewEncoder(w).Encode(response)
}

func handleDeleteImage(w http.ResponseWriter, id int) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "delete image by id " + strconv.Itoa(id),
	}

	json.NewEncoder(w).Encode(response)
}
