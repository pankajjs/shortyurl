package utils

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GetRandomUuid() string {
	return uuid.NewString()
}