package role

import (
	"base-api/internal/pkg/middleware"
	"net/http"
)

func (h *RoleHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle(
		"GET /api/roles",
		middleware.Auth(middleware.Permission("view:roles")(http.HandlerFunc(h.GetAll))),
	)
	mux.Handle(
		"GET /api/roles/{id}",
		middleware.Auth(middleware.Permission("view:roles")(http.HandlerFunc(h.GetByID))),
	)
	mux.Handle(
		"POST /api/roles",
		middleware.Auth(middleware.Permission("create:roles")(http.HandlerFunc(h.Create))),
	)
	mux.Handle(
		"PUT /api/roles/{id}",
		middleware.Auth(middleware.Permission("update:roles")(http.HandlerFunc(h.Update))),
	)
	mux.Handle(
		"DELETE /api/roles/{id}",
		middleware.Auth(middleware.Permission("delete:roles")(http.HandlerFunc(h.Delete))),
	)
}
