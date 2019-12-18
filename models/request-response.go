package models

type IsAdminRequest struct {
	Token string `json:"token"`
}

type IsAdminResponse struct {
	Admin bool  `json:"admin"`
	Err   error `json:"error,omitempty"`
}

type ChangePermissionRequest struct {
	Phone   string `json:"phone"`
	Promote bool   `json:"promote"`
}

type ChangePermissionResponse struct {
	Err error `json:"error,omitempty"`
}
