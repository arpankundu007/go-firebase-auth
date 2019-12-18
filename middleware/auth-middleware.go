package middleware

import (
	"context"
	"go-firebase-auth/utils"
	"net/http"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := utils.GetFirebaseAuthClient()
		if err != nil {
			utils.WriteJSON(w, err.Error())
			return
		}
		token, err := utils.GetIdTokenFromHeader(r)
		if err!=nil{
			utils.WriteJSON(w, err.Error())
			return
		}
		_, err = client.VerifyIDToken(context.Background(), token)
		if err != nil {
			utils.WriteJSON(w, err.Error())
			return
		}
		next.ServeHTTP(w, r)
	})


}
