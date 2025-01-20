package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"band-manager-backend/internal/usecases/helpers"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

// GroupUsecase implements group management logic.
type GroupUsecase struct {
	groupRepo *repositories.GroupRepository
	userRepo  *repositories.UserRepository
}

func NewGroupUsecase() *GroupUsecase {
	return &GroupUsecase{
		groupRepo: repositories.NewGroupRepository(),
		userRepo:  repositories.NewUserRepository(),
	}
}

// RefreshAccessToken generates and updates a new access token for a group.
func (u *GroupUsecase) RefreshAccessToken(groupID uint, requestingUserID uint) (string, error) {
	role, err := u.groupRepo.GetUserRole(requestingUserID, groupID)
	if err != nil {
		return "", errors.New("user not in group")
	}

	if role != helpers.RoleManager {
		return "", errors.New("insufficient permissions - only managers can refresh access token")
	}

	newToken := generateAccessToken()

	err = u.groupRepo.UpdateAccessToken(groupID, newToken)
	if err != nil {
		return "", fmt.Errorf("failed to update access token: %v", err)
	}

	return newToken, nil
}

// MemberInfo holds the details of a member in a group.
type MemberInfo struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

// GroupInfo holds the details of a group.
type GroupInfo struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Role         string `json:"role"`
	MembersCount int    `json:"members_count"`
}

// generateAccessToken generates a random access token for group access.
func generateAccessToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// CreateGroup creates a new band group and sets up initial roles.
func (u *GroupUsecase) CreateGroup(name, description string, userID uint) (string, uint, error) {
	_, err := u.userRepo.GetUserByID(userID)

	if err != nil {
		return "", 0, fmt.Errorf("User not found: %v", err)
	}

	group := &model.Group{
		Name:        name,
		Description: description,
		AccessToken: generateAccessToken(),
	}

	err = u.groupRepo.CreateGroup(group)
	if err != nil {
		return "", 0, fmt.Errorf("Could not create group: %v", err)
	}

	err = u.groupRepo.AddUserToGroup(userID, group.ID, helpers.RoleManager)
	if err != nil {
		return "", 0, fmt.Errorf("Could not add user to group: %v", err)
	}

	return helpers.RoleManager, group.ID, nil
}

// JoinGroup processes a user's request to join a group via access token.
func (u *GroupUsecase) JoinGroup(userID uint, accessToken string) (string, uint, string, error) {
	group, err := u.groupRepo.GetGroupByAccessToken(accessToken)
	if err != nil {
		return "", 0, "", errors.New("invalid access token")
	}

	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return "", 0, "", errors.New("user not found")
	}

	for _, g := range user.Groups {
		if g.ID == group.ID {
			return "", 0, "", errors.New("user already in group")
		}
	}

	err = u.groupRepo.AddUserToGroup(userID, group.ID, helpers.RoleMember)
	if err != nil {
		return "", 0, "", errors.New("failed to join group")
	}

	return helpers.RoleMember, group.ID, group.Name, nil
}

// GetGroupInfo retrieves group details and access token for managers.
func (u *GroupUsecase) GetGroupInfo(userID uint, groupID uint) (string, string, string, error) {
	role, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return "", "", "", errors.New("user not in group")
	}

	group, err := u.groupRepo.GetGroupByID(groupID)
	if err != nil {
		return "", "", "", errors.New("could not find group")
	}

	accessToken := ""
	if role == helpers.RoleManager {
		accessToken = group.AccessToken
	}

	return group.Name, group.Description, accessToken, nil
}

// GetGroupMembers retrieves all members of a group with their roles.
func (u *GroupUsecase) GetGroupMembers(groupID, requestingUserID uint) ([]MemberInfo, error) {
	_, err := u.groupRepo.GetUserRole(requestingUserID, groupID)
	if err != nil {
		return nil, errors.New("user not authorized to view this group")
	}

	members, err := u.groupRepo.GetGroupMembers(groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group members: %v", err)
	}

	var memberInfos []MemberInfo
	for _, member := range members {
		role, err := u.groupRepo.GetUserRole(member.ID, groupID)
		if err != nil {
			continue
		}
		memberInfos = append(memberInfos, MemberInfo{
			ID:        member.ID,
			FirstName: member.FirstName,
			LastName:  member.LastName,
			Email:     member.Email,
			Role:      role,
		})
	}

	return memberInfos, nil
}

// GetUserGroups retrieves all groups a user belongs to.
func (u *GroupUsecase) GetUserGroups(userID uint) ([]GroupInfo, error) {
	roles, err := u.userRepo.GetUserGroupRoles(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles: %v", err)
	}

	var groupInfos []GroupInfo
	for _, role := range roles {
		group, err := u.groupRepo.GetGroupByID(role.GroupID)
		if err != nil {
			continue
		}

		membersCount := len(group.Users)

		groupInfos = append(groupInfos, GroupInfo{
			ID:           group.ID,
			Name:         group.Name,
			Description:  group.Description,
			Role:         role.Role,
			MembersCount: membersCount,
		})
	}

	return groupInfos, nil
}

// RemoveMember removes a user from a group if requester has permissions.
func (u *GroupUsecase) RemoveMember(groupID, userToRemoveID, requestingUserID uint) error {

	requesterRole, err := u.groupRepo.GetUserRole(requestingUserID, groupID)
	if err != nil {
		return errors.New("requesting user not in group")
	}

	if !helpers.IsManagerOrModeratorRole(requesterRole) {
		return errors.New("insufficient permissions")
	}

	if userToRemoveID == requestingUserID {
		return errors.New("cannot remove yourself from group")
	}

	return u.groupRepo.RemoveUserFromGroup(userToRemoveID, groupID)
}

// UpdateMemberRole changes a user's role within a group.
func (u *GroupUsecase) UpdateMemberRole(groupID uint, userToUpdateID uint, requestingUserID uint, newRole string) error {
	requesterRole, err := u.groupRepo.GetUserRole(requestingUserID, groupID)
	if err != nil {
		return errors.New("requesting user not in group")
	}

	if requesterRole != helpers.RoleManager {
		return errors.New("insufficient permissions - only managers can change roles")
	}

	if userToUpdateID == requestingUserID {
		return errors.New("cannot change your own role")
	}

	if !isValidRole(newRole) {
		return errors.New("invalid role - must be 'manager', 'moderator', or 'member'")
	}

	return u.groupRepo.UpdateUserRole(userToUpdateID, groupID, newRole)
}

// isValidRole checks if the provided role is valid.
func isValidRole(role string) bool {
	validRoles := []string{helpers.RoleManager, helpers.RoleModerator, helpers.RoleMember}
	for _, r := range validRoles {
		if r == role {
			return true
		}
	}
	return false
}
