package models

type Request struct {
	Header
	RequestBody
}

type RequestBody struct {
	Body interface{}
}

type Header struct {
	ContentType string `json:"content-type"`
	AppVersion string`json:"app-version"`
	Token string `json:"token"`
}

type IsAdminResponse struct {
	Admin bool  `json:"admin"`
	Err   error `json:"error,omitempty"`
}

type ChangePermissionRequest struct {
	Header
	Phone   string `json:"phone"`
	Promote bool   `json:"promote"`
}

type ChangePermissionResponse struct {
	Err error `json:"error,omitempty"`
}

type WeatherResponse struct {
	Err error `json:"error,omitempty"`
}

