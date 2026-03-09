package server

import (
	"base-api/internal/handler/auth"
	"base-api/internal/handler/config"
	"base-api/internal/handler/permission"

	"base-api/internal/handler/role"
	"base-api/internal/handler/user"
	"base-api/internal/pkg/middleware"
	"base-api/internal/repository"
	"base-api/internal/service"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	HttpServer *http.Server
}

func NewServer(mux *http.ServeMux, db *sqlx.DB) *Server {

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, db)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(mux)

	authService := service.NewAuthService(userRepo)
	authHandler := auth.NewHandler(authService, *validate)
	authHandler.RegisterRoutes(mux)

	roleRepo := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepo)
	roleHandler := role.NewRoleHandler(roleService)
	roleHandler.RegisterRoutes(mux)

	permissionRepo := repository.NewPermissionRepository(db)
	permissionService := service.NewPermissionService(permissionRepo)
	permissionHandler := permission.NewPermissionHandler(permissionService)
	permissionHandler.RegisterRoutes(mux)



	configRepo := repository.NewConfigRepository(db)
	configService := service.NewConfigService(configRepo)
	configHandler := config.NewConfigHandler(configService)
	configHandler.RegisterRoutes(mux)

	configValueRepo := repository.NewConfigValueRepository(db)
	configValueService := service.NewConfigValueService(configValueRepo)
	configValueHandler := config.NewConfigValueHandler(configValueService)
	configValueHandler.RegisterRoutes(mux)



	handler := middleware.CORSMiddleware(mux)
	handler = middleware.Logging(handler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	return &Server{
		HttpServer: s,
	}
}
