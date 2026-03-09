package role

import (
	dto "base-api/internal/dto/role"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/service"
	"net/http"

	"github.com/google/uuid"
)

type RoleHandler struct {
	service service.RoleService
}

func NewRoleHandler(service service.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

func (h *RoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	roles, err := h.service.GetAll(r.Context())
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR)
		return
	}

	response := map[string]interface{}{
		"data":    roles,
		"message": "Get roles successfully",
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *RoleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		helper.WriteError(w, http.StatusBadRequest, constants.ID_IS_REQUIRED)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	role, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR)
		return
	}

	response := map[string]interface{}{
		"data":    role,
		"message": "Get role successfully",
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.RoleCreateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(r.Context(), req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Role created successfully"})
}

func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	var req dto.RoleUpdateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Update(r.Context(), id, req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Role updated successfully"})
}

func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Role deleted successfully"})
}
