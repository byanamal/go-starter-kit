package user

import (
	"base-api/internal/pkg/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle(
		"GET /api/users",
		middleware.Auth(middleware.Permission("view:users")(http.HandlerFunc(h.GetAll))),
	)
	mux.Handle(
		"GET /api/users/{id}",
		middleware.Auth(middleware.Permission("view:users")(http.HandlerFunc(h.GetByID))),
	)
	mux.Handle(
		"POST /api/users",
		middleware.Auth(middleware.Permission("create:users")(http.HandlerFunc(h.Create))),
	)
	mux.Handle(
		"PUT /api/users/{id}",
		middleware.Auth(middleware.Permission("update:users")(http.HandlerFunc(h.Update))),
	)
	mux.Handle(
		"DELETE /api/users/{id}",
		middleware.Auth(middleware.Permission("delete:users")(http.HandlerFunc(h.Delete))),
	)
}
