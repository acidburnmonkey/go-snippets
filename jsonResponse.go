package main

// pkg/response/json.go
// package response

import (
	"encoding/json"
	"net/http"
)

// JSON() -> sends a json response with status code + data \n
// ex: response.JSON(w, http.StatusOK, events)
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Error() -> sends a json response error + message
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}
