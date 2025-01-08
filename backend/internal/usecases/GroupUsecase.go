package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

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

type MemberInfo struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type GroupInfo struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Role         string `json:"role"`
	MembersCount int    `json:"members_count"`
}

// generateAccessToken creates a random hex string for group access
func generateAccessToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (u *GroupUsecase) CreateGroup(name, description string, userID uint) (string, uint, error) {
	_, err := u.userRepo.GetUserByID(userID)

	if err != nil {
		return "", 0, fmt.Errorf("użytkownik nie znaleziony: %v", err)
	}

	group := &model.Group{
		Name:        name,
		Description: description,
		AccessToken: generateAccessToken(),
	}

	err = u.groupRepo.CreateGroup(group)
	if err != nil {
		return "", 0, fmt.Errorf("nie udało się utworzyć grupy: %v", err)
	}

	err = u.groupRepo.AddUserToGroup(userID, group.ID, "manager")
	if err != nil {
		return "", 0, fmt.Errorf("nie udało się dodać użytkownika do grupy: %v", err)
	}

	return "manager", group.ID, nil
}

func (u *GroupUsecase) JoinGroup(userID uint, accessToken string) (string, uint, error) {
	group, err := u.groupRepo.GetGroupByAccessToken(accessToken)
	if err != nil {
		return "", 0, errors.New("invalid access token")
	}

	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return "", 0, errors.New("user not found")
	}

	for _, g := range user.Groups {
		if g.ID == group.ID {
			return "", 0, errors.New("user already in group")
		}
	}

	err = u.groupRepo.AddUserToGroup(userID, group.ID, "member")
	if err != nil {
		return "", 0, errors.New("failed to join group")
	}

	return "member", group.ID, nil
}

func (u *GroupUsecase) GetGroupInfo(userID uint, groupID uint) (string, string, string, error) {
	role, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return "", "", "", fmt.Errorf("użytkownik nie należy do tej grupy")
	}

	group, err := u.groupRepo.GetGroupByID(groupID)
	if err != nil {
		return "", "", "", fmt.Errorf("nie znaleziono grupy")
	}

	accessToken := ""
	if role == "manager" {
		accessToken = group.AccessToken
	}

	return group.Name, group.Description, accessToken, nil
}

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

func (u *GroupUsecase) RemoveMember(groupID, userToRemoveID, requestingUserID uint) error {

	requesterRole, err := u.groupRepo.GetUserRole(requestingUserID, groupID)
	if err != nil {
		return errors.New("requesting user not in group")
	}

	if requesterRole != "manager" && requesterRole != "moderator" {
		return errors.New("insufficient permissions")
	}

	if userToRemoveID == requestingUserID {
		return errors.New("cannot remove yourself from group")
	}

	return u.groupRepo.RemoveUserFromGroup(userToRemoveID, groupID)
}
