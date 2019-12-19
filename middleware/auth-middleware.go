package middleware

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go-firebase-auth/models"
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

func GoKitIsAuthorised() endpoint.Middleware{
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, r interface{}) (response interface{}, err error) {
			request := r.(models.Request)
			fmt.Println("ContentType: ", request.ContentType)
			fmt.Println("App-Version: ", request.AppVersion)
			fmt.Println("Authorization: ", request.Token)
			client, err := utils.GetFirebaseAuthClient()
			if err != nil {
				return err.Error(), err
			}
			_, err = client.VerifyIDToken(context.Background(), request.Token)
			if err != nil {
				return err.Error(), err
			}
			return next(ctx, r)
		}
	}
}

func HttpMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodPost{
			next.ServeHTTP(w, r)
		}else{
			utils.WriteJSON(w, "Invalid HTTP Method")
		}
	})
}

func GoKitMW(next endpoint.Endpoint) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return next(ctx, request)
	}
}

