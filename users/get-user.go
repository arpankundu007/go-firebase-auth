package users

import (
	"context"
	auth2 "firebase.google.com/go/auth"
	"fmt"
	"go-firebase-auth/auth"
	"go-firebase-auth/utils"
	"net/http"
)

func GetUserFromPhone() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		phone := utils.GetParamFromRequestUrl(r, "phone")
		fmt.Println(phone)
		user, err := getUserFromPhone(phone)
		if err!=nil{
			utils.WriteJSON(w, err.Error())
			return
		}
		utils.WriteJSON(w, user)
	})
}

func getUserFromPhone(phone string) (*auth2.UserRecord, error){
	client, err := auth.GetFirebaseAuthClient()
	ctx := context.Background()
	if err!=nil{
		return nil, err
	}
	u, err := client.GetUserByPhoneNumber(ctx, phone)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserIdFromPhoneNumber(phone string) (string, error) {
	client, err := auth.GetFirebaseAuthClient()
	ctx := context.Background()
	u, err := client.GetUserByPhoneNumber(ctx, phone)
	if err != nil {
		return "nil", err
	}
	return u.UID, nil
}
