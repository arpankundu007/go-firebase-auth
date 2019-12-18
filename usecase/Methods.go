package usecase

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-firebase-auth/utils"
	"io/ioutil"
	"net/http"
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

func (AuthServiceInstance) ChangePermission(phone string, promote bool) error{
	client, err := utils.GetFirebaseAuthClient()
	if err != nil {
		return err
	}
	uid, err := utils.GetUserIdFromPhoneNumber(phone)
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

func (AuthServiceInstance) GetWeather() endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		url := "https://samples.openweathermap.org/data/2.5/forecast?q=M%C3%BCnchen,DE&appid=b6907d289e10d714a6e88b30761fae22"

		req, err := http.NewRequest("GET", url, nil)
		if err!=nil{
			return err.Error(), err
		}

		res, err := http.DefaultClient.Do(req)
		if err!=nil{
			return err.Error(), err
		}

		body, err := ioutil.ReadAll(res.Body)
		if err!=nil{
			return err.Error(), err
		}

		return string(body), nil
	}
}


