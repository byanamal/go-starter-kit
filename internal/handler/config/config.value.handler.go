package config

import (
	dto "base-api/internal/dto/config"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/service"
	"net/http"

	"github.com/google/uuid"
)

type ConfigValueHandler struct {
	service service.ConfigValueService
}

func NewConfigValueHandler(service service.ConfigValueService) *ConfigValueHandler {
	return &ConfigValueHandler{service: service}
}

func (h *ConfigValueHandler) GetValuesByConfigID(w http.ResponseWriter, r *http.Request) {
	configIDStr := r.PathValue("id")
	if configIDStr == "" {
		helper.WriteError(w, http.StatusBadRequest, constants.ID_IS_REQUIRED)
		return
	}

	configID, err := uuid.Parse(configIDStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	values, err := h.service.GetValuesByConfigID(r.Context(), configID)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR)
		return
	}

	response := map[string]interface{}{
		"data":    values,
		"message": "Get config values successfully",
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *ConfigValueHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	val, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR)
		return
	}

	response := map[string]interface{}{
		"data":    val,
		"message": "Get config value successfully",
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *ConfigValueHandler) Create(w http.ResponseWriter, r *http.Request) {
	configIDStr := r.PathValue("configId")
	configID, err := uuid.Parse(configIDStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	var req dto.ConfigValueCreateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(r.Context(), configID, req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Config value created successfully"})
}

func (h *ConfigValueHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	var req dto.ConfigValueUpdateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Update(r.Context(), id, req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Config value updated successfully"})
}

func (h *ConfigValueHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Config value deleted successfully"})
}
