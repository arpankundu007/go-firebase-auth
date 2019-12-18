package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-firebase-auth/models"
	"go-firebase-auth/usecase"
)

func GetIsAdminEndpoint(svc usecase.Service) endpoint.Endpoint{
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.IsAdminRequest)

		admin, err := svc.IsAdmin(req.Token)

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
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.ChangePermissionRequest)

		permError := svc.ChangePermission(req.Phone, req.Promote)

		if permError!=nil{
			return models.ChangePermissionResponse{Err: permError}, permError
		}
		return models.ChangePermissionResponse{Err: nil}, nil
	}
}
