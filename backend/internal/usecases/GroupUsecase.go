package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"crypto/rand"
	"encoding/hex"
	"errors"
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
	// Create new group with access token
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

	// Update creator's role and group
	creator, err := u.userRepo.GetUserByID(creatorID)
	if err != nil {
		return "", 0, errors.New("failed to find creator user")
	}

	creator.GroupID = &group.ID
	creator.Role = "admin"

	err = u.userRepo.UpdateUser(creator)
	if err != nil {
		return "", 0, errors.New("failed to update creator's details")
	}

	return creator.Role, group.ID, nil
}
