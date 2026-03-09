package config

import (
	"base-api/internal/pkg/middleware"
	"net/http"
)

func (h *ConfigHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle(
		"GET /api/configs",
		middleware.Auth(middleware.Permission("view:configs")(http.HandlerFunc(h.GetAll))),
	)
	mux.Handle(
		"GET /api/configs/{id}",
		middleware.Auth(middleware.Permission("view:configs")(http.HandlerFunc(h.GetByID))),
	)
	mux.Handle(
		"POST /api/configs",
		middleware.Auth(middleware.Permission("create:configs")(http.HandlerFunc(h.Create))),
	)
	mux.Handle(
		"PUT /api/configs/{id}",
		middleware.Auth(middleware.Permission("update:configs")(http.HandlerFunc(h.Update))),
	)
	mux.Handle(
		"DELETE /api/configs/{id}",
		middleware.Auth(middleware.Permission("delete:configs")(http.HandlerFunc(h.Delete))),
	)
}
