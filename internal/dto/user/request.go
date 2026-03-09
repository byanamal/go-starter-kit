package dto

type UserCreateRequest struct {
	Email    string  `json:"email" validate:"required,email"`
	Name     *string `json:"name" validate:"required,min=3,max=100"`
	Password string  `json:"password" validate:"required,min=8"`
}

type UserUpdateRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	Name     *string `json:"name" validate:"omitempty,min=3,max=100"`
	Password *string `json:"password" validate:"omitempty,min=8"`
}
