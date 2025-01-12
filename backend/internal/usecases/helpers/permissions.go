package helpers

const (
	RoleManager   = "manager"
	RoleMember    = "member"
	RoleModerator = "moderator"
)

func HasManagerOrModeratorRole(role string) bool {
	return role == RoleManager || role == RoleModerator
}
