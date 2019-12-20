package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"go-firebase-auth/utils"
	"net/http"
	"strings"
)

// net/http Middleware
func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := utils.GetFirebaseAuthClient()
		if err != nil {
			utils.WriteJSON(w, err.Error())
			return
		}
		token, err := utils.GetIdTokenFromHeader(r)
		if err != nil {
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

//Go-Kit middleware
func GoKitIsAuthorised(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (response interface{}, err error) {
		token := strings.Split(ctx.Value(httptransport.ContextKeyRequestAuthorization).(string), " ")[1]
		client, err := utils.GetFirebaseAuthClient()
		if err != nil {
			return err.Error(), err
		}
		_, err = client.VerifyIDToken(context.Background(), token)
		if err != nil {
			return err.Error(), err
		}
		return next(ctx, r)
	}

}
