package dto

type PermissionCreateRequest struct {
	Name  string `json:"name" validate:"required"`
	Code  string `json:"code" validate:"required"`
	Group string `json:"group"`
}

type PermissionUpdateRequest struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Group string `json:"group"`
}
