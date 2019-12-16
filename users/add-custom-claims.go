package users

import (
	"context"
	"firebase_auth/auth"
	"firebase_auth/utils"
	"net/http"
)

func PromoteUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := modifyUserLever(utils.GetParamFromRequestUrl(r, "phone"), true)
		if err != nil {
			utils.WriteJSON(w, err.Error())
		}
	})
}

func DemoteUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := modifyUserLever(utils.GetParamFromRequestUrl(r, "phone"), false)
		if err != nil {
			utils.WriteJSON(w, err.Error())
		}
	})
}

func modifyUserLever(phone string, promote bool) error {
	client, err := auth.GetFirebaseAuthClient()
	if err != nil {
		return err
	}
	uid, err := GetUserIdFromPhoneNumber(phone)
	if err != nil {
		return err
	}
	if promote {
		claims := map[string]interface{}{"admin": true}
		err = client.SetCustomUserClaims(context.Background(), uid, claims)
	}else{
		claims := map[string]interface{}{"admin": false}
		err = client.SetCustomUserClaims(context.Background(), uid, claims)
	}
	if err != nil {
		return err
	}
	return nil
}