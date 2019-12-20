package usecase

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-firebase-auth/utils"
	"strings"
)

type AuthServiceInstance struct{}

func (AuthServiceInstance) IsAdmin(token string) (bool, error){
	client, _ := utils.GetFirebaseAuthClient()
	verifiedToken, err := client.VerifyIDToken(context.Background(), token)

	if err != nil {
		return false, err
	}

	claims := verifiedToken.Claims
	if admin, ok := claims["admin"]; ok {
		if admin.(bool) {
			return true, nil
		}else {
			return false, nil
		}
	}else{
		return false, errors.New("claim does not exist")
	}
}

func (AuthServiceInstance) ChangePermission(phone string, admin string) error{
	client, err := utils.GetFirebaseAuthClient()
	if err != nil {
		return err
	}
	uid, err := utils.GetUserIdFromPhoneNumber(phone)
	if err != nil {
		return err
	}
	if strings.Compare(admin, "promote") == 0 {
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

func (AuthServiceInstance) HelloWorld() endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return "Hello World", nil
	}
}


