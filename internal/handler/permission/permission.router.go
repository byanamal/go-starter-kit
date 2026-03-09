package permission

import (
	"base-api/internal/pkg/middleware"
	"net/http"
)

func (h *PermissionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle(
		"GET /api/permissions",
		middleware.Auth(middleware.Permission("view:permissions")(http.HandlerFunc(h.GetAll))),
	)
	mux.Handle(
		"GET /api/permissions/{id}",
		middleware.Auth(middleware.Permission("view:permissions")(http.HandlerFunc(h.GetByID))),
	)
	mux.Handle(
		"POST /api/permissions",
		middleware.Auth(middleware.Permission("create:permissions")(http.HandlerFunc(h.Create))),
	)
	mux.Handle(
		"PUT /api/permissions/{id}",
		middleware.Auth(middleware.Permission("update:permissions")(http.HandlerFunc(h.Update))),
	)
	mux.Handle(
		"DELETE /api/permissions/{id}",
		middleware.Auth(middleware.Permission("delete:permissions")(http.HandlerFunc(h.Delete))),
	)
}
