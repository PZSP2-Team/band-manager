package helpers

const (
	RoleManager   = "manager"
	RoleMember    = "member"
	RoleModerator = "moderator"
)

func IsManagerOrModeratorRole(role string) bool {
	return role == RoleManager || role == RoleModerator
}
