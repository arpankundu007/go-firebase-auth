package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"go-firebase-auth/models"
	"go-firebase-auth/usecase"
	"strings"
)

func GetIsAdminEndpoint(svc usecase.Service) endpoint.Endpoint{
	return func(context context.Context, request interface{}) (interface{}, error) {
		authHeader := context.Value(httptransport.ContextKeyRequestAuthorization).(string)
		token := strings.Split(authHeader, " ")[1]
		admin, err := svc.IsAdmin(token)

		if err!=nil{
			return models.IsAdminResponse{
				Admin: false,
				Err:   err,
			}, err
		}
		return models.IsAdminResponse{
			Admin: admin,
			Err:   nil,
		}, nil
	}
}

func GetChangePermissionEndpoint(svc usecase.Service) endpoint.Endpoint{
	return func(context context.Context, request interface{}) (interface{}, error) {

		requestBody := request.(models.ChangePermissionRequest)

		permError := svc.ChangePermission(requestBody.Phone, requestBody.Promote)

		if permError!=nil{
			return models.ChangePermissionResponse{Err: permError, Message:permError.Error()}, permError
		}
		return models.ChangePermissionResponse{Err: nil, Message:"Successful"}, nil
	}
}
