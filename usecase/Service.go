package usecase

import "github.com/go-kit/kit/endpoint"

type Service interface {
	IsAdmin(token string) (bool, error)
	ChangePermission(phone string, promote bool) error
	GetWeather() endpoint.Endpoint
}