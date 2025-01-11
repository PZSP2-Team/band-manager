package helpers

const (
	RoleManager   = "manager"
	RoleModerator = "moderator"
)

func HasManagerOrModeratorRole(role string) bool {
	return role == RoleManager || role == RoleModerator
}
