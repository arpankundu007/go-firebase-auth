package transport

import (
	"context"
	"encoding/json"
	"go-firebase-auth/models"
	"net/http"
)

//Decoders parse the HTTP Requests and returns them as interface{}
func IsAdminRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return "", nil
}

func ChangePermissionRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var cpr models.ChangePermissionRequest
	err := json.NewDecoder(r.Body).Decode(&cpr)
	if err!=nil{
		return nil, err
	}
	return cpr, nil
}

func GetHelloWorldDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return "hello, world", nil
}
