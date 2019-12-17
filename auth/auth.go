package auth

import (
	"context"
	"errors"
	"firebase.google.com/go/auth"
	"fmt"
	"go-firebase-auth/firebase"
	"go-firebase-auth/utils"
	"log"
	"net/http"
	"strings"
)

func IsAuthorised(next http.Handler) http.Handler {
	app, err := firebase.GetFirebaseApp()
	ctx := context.Background()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client, err := app.Auth(ctx)
			if err != nil {
				utils.WriteJSON(w, "error getting Auth client: " + err.Error())
				return
			}
			headerToken, err := getIdTokenFromHeader(r)
			if err != nil {
				utils.WriteJSON(w, err.Error())
				return
			}
			token, err := client.VerifyIDToken(ctx, headerToken)
			if err != nil {
				utils.WriteJSON(w, "error verifying ID token: "+ err.Error())
				return
			}
			next.ServeHTTP(w, r)
			log.Printf("Verified ID token: %v\n", token)
		})
	}
}

func GetFirebaseAuthClient() (*auth.Client, error){
	app, err := firebase.GetFirebaseApp()
	ctx := context.Background()

	if err!=nil{
		return nil, err
	}

	client, err := app.Auth(ctx)
	return client, err
}

func getIdTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) != 2 {
		return "", errors.New("invalid auth header")
	}
	return splitAuthHeader[1], nil
}
