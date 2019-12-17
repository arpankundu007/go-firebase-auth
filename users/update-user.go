package users

import (
	"context"
	"firebase.google.com/go/auth"
	auth2 "go-firebase-auth/auth"
	"go-firebase-auth/models"
	"go-firebase-auth/utils"
	"net/http"
)

func UpdateUser() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		user, err := utils.ReadJSON(r)
		if err!=nil{
			utils.WriteJSON(w, err.Error())
			return
		}
		phone := utils.GetParamFromRequestUrl(r, "phone")
		_, err = updateUserByPhoneNumber(phone, user)
		if err!=nil{
			utils.WriteJSON(w, err.Error())
			return
		}
	})
}

func updateUserByPhoneNumber(phone string, user models.User) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		DisplayName(user.Name)
	client, err := auth2.GetFirebaseAuthClient()
	ctx := context.Background()
	uid, err := GetUserIdFromPhoneNumber(phone)
	if err!=nil{
		return nil, err
	}
	u, err := client.UpdateUser(ctx, uid, params)
	if err != nil {
		return nil, err
	}
	return u, nil
}
