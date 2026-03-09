package config

import (
	dto "base-api/internal/dto/config"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/service"
	"net/http"

	"github.com/google/uuid"
)

type ConfigHandler struct {
	service service.ConfigService
}

func NewConfigHandler(service service.ConfigService) *ConfigHandler {
	return &ConfigHandler{service: service}
}

// Config Handlers

func (h *ConfigHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	configs, err := h.service.GetAll(r.Context())
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR)
		return
	}

	response := map[string]interface{}{
		"data":    configs,
		"message": "Get all configs successfully",
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *ConfigHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	cfg, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		helper.WriteError(w, http.StatusNotFound, "Config not found")
		return
	}

	response := map[string]interface{}{
		"data":    cfg,
		"message": "Get config successfully",
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *ConfigHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.ConfigCreateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(r.Context(), req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Config created successfully"})
}

func (h *ConfigHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	var req dto.ConfigUpdateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Update(r.Context(), id, req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Config updated successfully"})
}

func (h *ConfigHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Config deleted successfully"})
}
