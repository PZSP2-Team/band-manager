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
