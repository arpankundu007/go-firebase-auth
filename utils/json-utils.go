package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func WriteJSON(w http.ResponseWriter, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	parseError := json.NewEncoder(w).Encode(&data)
	if parseError != nil {
		panic(parseError.Error())
	}
}

func GetIdTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) != 2 {
		return "", errors.New("invalid auth header")
	}
	return splitAuthHeader[1], nil
}
