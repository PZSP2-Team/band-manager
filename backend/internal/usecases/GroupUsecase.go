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

func (u *GroupUsecase) CreateGroup(name, description string, creatorID uint) (string, uint, error) {

	accessToken := generateAccessToken()
	group := &model.Group{
		Name:        name,
		Description: description,
		AccessToken: accessToken,
	}

	err := u.groupRepo.CreateGroup(group)
	if err != nil {
		return "", 0, errors.New("failed to create group")
	}

	creator, err := u.userRepo.GetUserByID(creatorID)
	if err != nil {
		return "", 0, errors.New("failed to find creator user")
	}

	creator.GroupID = &group.ID
	creator.Role = "manager"

	err = u.userRepo.UpdateUser(creator)
	if err != nil {
		return "", 0, errors.New("failed to update creator's details")
	}

	return creator.Role, group.ID, nil
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

	if user.GroupID != nil {
		return "", 0, errors.New("user already belongs to a group")
	}

	user.GroupID = &group.ID
	user.Role = "member"

	err = u.userRepo.UpdateUser(user)
	if err != nil {
		return "", 0, errors.New("failed to join group")
	}

	return user.Role, group.ID, nil
}

func (u *GroupUsecase) GetGroupInfo(userID uint) (string, string, string, error) {

	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return "", "", "", fmt.Errorf("użytkownik nie znaleziony")
	}

	if user.GroupID == nil {
		return "", "", "", fmt.Errorf("użytkownik nie należy do żadnej grupy")
	}

	group, err := u.groupRepo.GetGroupByID(*user.GroupID)
	if err != nil {
		return "", "", "", fmt.Errorf("nie znaleziono grupy")
	}

	accessToken := ""
	if user.Role == "manager" {
		accessToken = group.AccessToken
	}

	return group.Name, group.Description, accessToken, nil
}
