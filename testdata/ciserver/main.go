package main

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api/v1/health1", health1Handler)
	http.HandleFunc("/api/v1/health2", health2Handler)
	http.HandleFunc("/api/v1/health3", health3Handler)
	http.HandleFunc("/api/v1/health4", health4Handler)

	http.ListenAndServe(":8080", nil)
}

func health1Handler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "ok", Message: "Service 1 is healthy"}
	json.NewEncoder(w).Encode(response)
}

func health2Handler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "ok", Message: "Service 2 is healthy"}
	json.NewEncoder(w).Encode(response)
}

func health3Handler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "ok", Message: "Service 3 is healthy"}
	json.NewEncoder(w).Encode(response)
}

func health4Handler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "ok", Message: "Service 4 is healthy"}
	json.NewEncoder(w).Encode(response)
}
