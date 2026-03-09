package config

import (
	"base-api/internal/pkg/middleware"
	"net/http"
)

func (h *ConfigValueHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle(
		"GET /api/config/{configId}/values",
		middleware.Auth(middleware.Permission("view:config-values")(http.HandlerFunc(h.GetValuesByConfigID))),
	)
	mux.Handle(
		"GET /api/config-values/{id}",
		middleware.Auth(middleware.Permission("view:config-values")(http.HandlerFunc(h.GetByID))),
	)
	mux.Handle(
		"POST /api/config/{configId}/values",
		middleware.Auth(middleware.Permission("create:config-values")(http.HandlerFunc(h.Create))),
	)
	mux.Handle(
		"PUT /api/config-values/{id}",
		middleware.Auth(middleware.Permission("update:config-values")(http.HandlerFunc(h.Update))),
	)
	mux.Handle(
		"DELETE /api/config-values/{id}",
		middleware.Auth(middleware.Permission("delete:config-values")(http.HandlerFunc(h.Delete))),
	)
}
