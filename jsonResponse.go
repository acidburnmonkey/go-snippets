package main

// package response

import (
	"encoding/json"
	"net/http"
)

// JsonResponse() -> sends a json response with status code + data \n
// ex: JsonResponse(w, http.StatusOK, events)
func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ErrorResponse() -> sends a json response error + message
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	JsonResponse(w, status, map[string]string{"error": message})
}
