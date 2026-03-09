package dto

type RoleCreateRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type RoleUpdateRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
