package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseModel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func statusJson(code int) ResponseModel {
	return ResponseModel{
		Status:  code,
		Message: http.StatusText(code),
	}
}

func encodeStatus(w http.ResponseWriter, status int) {
	json.NewEncoder(w).Encode(statusJson(status))
}

func handle(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func HandleOk(w http.ResponseWriter, r *http.Request) {
	handle(w, http.StatusOK)
	encodeStatus(w, http.StatusOK)
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	encodeStatus(w, http.StatusMethodNotAllowed)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	handle(w, http.StatusNotFound)
	encodeStatus(w, http.StatusNotFound)
}

func ResponseOK(w http.ResponseWriter) {
	encodeStatus(w, http.StatusOK)
}

func ResponseInternalServerError(w http.ResponseWriter) {
	encodeStatus(w, http.StatusInternalServerError)
}
