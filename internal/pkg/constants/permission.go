package constants

import (
	"fmt"
)

const (
	// Modules
	ModuleUser       = "users"
	ModuleRole       = "roles"
	ModulePermission = "permissions"
	ModuleMenu       = "menus"
	ModuleConfig     = "config"

	// Role Codes
	RoleCodeSuperadmin = "SUPERADMIN"
	RoleCodeAdmin      = "ADMIN"
	RoleCodeUser       = "USER"

	// Actions
	ActionView   = "view"
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionDelete = "delete"
)

func GetPermissionCode(module, action string) string {
	return fmt.Sprintf("%s:%s", action, module)
}

func GetPermissionName(module, action string) string {
	return fmt.Sprintf("%s %s", TitleCase(action), TitleCase(module))
}

func TitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

var (
	Modules = []string{
		ModuleUser,
		ModuleRole,
		ModulePermission,
		ModuleMenu,
		ModuleConfig,
	}
	Actions = []string{
		ActionView,
		ActionCreate,
		ActionUpdate,
		ActionDelete,
	}
)
