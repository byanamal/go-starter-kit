package permission

import (
	dto "base-api/internal/dto/permission"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/service"
	"net/http"

	"github.com/google/uuid"
)

type PermissionHandler struct {
	service service.PermissionService
}

func NewPermissionHandler(service service.PermissionService) *PermissionHandler {
	return &PermissionHandler{service: service}
}

func (h *PermissionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.service.GetAll(r.Context())
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR)
		return
	}

	response := map[string]interface{}{
		"data":    permissions,
		"message": "Get permissions successfully",
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *PermissionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	permission, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR)
		return
	}

	response := map[string]interface{}{
		"data":    permission,
		"message": "Get permission successfully",
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *PermissionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.PermissionCreateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(r.Context(), req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Permission created successfully"})
}

func (h *PermissionHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	var req dto.PermissionUpdateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Update(r.Context(), id, req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Permission updated successfully"})
}

func (h *PermissionHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Permission deleted successfully"})
}
