package transport

import (
	"context"
	"encoding/json"
	"go-firebase-auth/models"
	"go-firebase-auth/utils"
	"net/http"
	"strings"
)

func IsAdminRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	token, err := utils.GetIdTokenFromHeader(r)
	if err != nil {
		return nil, err
	}
	return models.IsAdminRequest{Token: token}, nil
}

func ChangePermissionRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	phone := utils.GetParamFromRequestUrl(r, "phone")
	admin := utils.GetParamFromRequestUrl(r, "admin")
	var promote bool
	if strings.Compare(admin, "promote") == 0 {
		promote = true
	} else if strings.Compare(admin, "demote") == 0 {
		promote = false
	}
	return models.ChangePermissionRequest{
		Phone:   phone,
		Promote: promote,
	}, nil
}

func GetWeatherDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request models.IsAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
