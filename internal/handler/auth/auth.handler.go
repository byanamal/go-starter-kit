package auth

import (
	dto "base-api/internal/dto/auth"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/pkg/validation"
	"base-api/internal/service"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  service.AuthService
	validate validator.Validate
}

func NewHandler(
	s service.AuthService,
	v validator.Validate,
) *Handler {
	return &Handler{
		service:  s,
		validate: v,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INVALID_REQUEST)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		helper.WriteValidationError(w, http.StatusBadRequest, validation.FormatValidationError(err))
		return
	}

	res, err := h.service.Login(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrInvalidCredential):
			helper.WriteError(w, http.StatusUnauthorized, constants.INVALID_CREDENTIAL)
		default:
			helper.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := map[string]interface{}{
		"data":    res,
		"message": "success",
	}

	helper.WriteJSON(w, http.StatusOK, response)

}
