package user

import (
	dto "base-api/internal/dto/user"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/service"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	service service.UserService
}

func NewHandler(
	s service.UserService,

) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	p := helper.GetPagination(r)

	data, total, err := h.service.GetAll(r.Context(), p.Page, p.Limit)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := helper.PaginationResult{
		Data: data,
		Meta: helper.PaginationMeta{
			Page:  p.Page,
			Limit: p.Limit,
			Total: total,
		},
	}

	helper.WriteJSON(w, http.StatusOK, res)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR)
		return
	}

	response := map[string]interface{}{
		"data": res,
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.UserCreateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(r.Context(), req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, constants.INVALID_FORMAT_UUID)
		return
	}

	var req dto.UserUpdateRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Update(r.Context(), id, req); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "User updated successfully"})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
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

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
