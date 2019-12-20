package models


type IsAdminResponse struct {
	Admin bool  `json:"admin"`
	Err   error `json:"error,omitempty"`
}

type ChangePermissionRequest struct {
	Phone   string `json:"phone"`
	Promote string   `json:"promote"`
}

type ChangePermissionResponse struct {
	Err error `json:"error,omitempty"`
}

type WeatherResponse struct {
	Err error `json:"error,omitempty"`
}

