package usecase

import "github.com/go-kit/kit/endpoint"

type Service interface {
	IsAdmin(token string) (bool, error)
	ChangePermission(phone string, promote string) error
	HelloWorld() endpoint.Endpoint
}