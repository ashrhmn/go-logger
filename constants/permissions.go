package constants

const (
	PermissionAddUser          = "ADD_USER"
	PermissionModifyUser       = "MODIFY_USER"
	PermissionViewUsers        = "VIEW_USERS"
	PermissionModifyPermission = "MODIFY_PERMISSION"
	PermissionAdmin            = "ADMIN"
	PermissionViewLogs         = "VIEW_LOGS"
)

var PermissionsAll = []string{
	PermissionAddUser,
	PermissionModifyUser,
	PermissionViewUsers,
	PermissionModifyPermission,
	PermissionAdmin,
	PermissionViewLogs,
}
