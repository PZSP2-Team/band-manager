package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"band-manager-backend/internal/usecases/helpers"
	"errors"
)

// SubgroupUsecase implements subgroup management logic.
type SubgroupUsecase struct {
	subgroupRepo *repositories.SubgroupRepository
	groupRepo    *repositories.GroupRepository
}

func NewSubgroupUsecase() *SubgroupUsecase {
	return &SubgroupUsecase{
		subgroupRepo: repositories.NewSubgroupRepository(),
		groupRepo:    repositories.NewGroupRepository(),
	}
}

// CreateSubgroup creates a new subgroup within a band group.
func (u *SubgroupUsecase) CreateSubgroup(name, description string, groupID uint, userID uint) (*model.Subgroup, error) {
	role, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}

	if !helpers.IsManagerOrModeratorRole(role) {
		return nil, errors.New("insufficient permissions")
	}

	subgroup := &model.Subgroup{
		Name:        name,
		Description: description,
		GroupID:     groupID,
	}

	if err := u.subgroupRepo.CreateSubgroup(subgroup); err != nil {
		return nil, err
	}

	return subgroup, nil
}

// GetSubgroup retrieves subgroup details if user has access.
func (u *SubgroupUsecase) GetSubgroup(id uint, userID uint) (*model.Subgroup, error) {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(id)
	if err != nil {
		return nil, err
	}

	_, err = u.groupRepo.GetUserRole(userID, subgroup.GroupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return subgroup, nil
}

// UpdateSubgroup modifies subgroup details if user has permissions.
func (u *SubgroupUsecase) UpdateSubgroup(id uint, name, description string, userID uint) error {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, subgroup.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if !helpers.IsManagerOrModeratorRole(role) {
		return errors.New("insufficient permissions")
	}

	subgroup.Name = name
	subgroup.Description = description

	return u.subgroupRepo.UpdateSubgroup(subgroup)
}

// DeleteSubgroup removes a subgroup if user has permissions.
func (u *SubgroupUsecase) DeleteSubgroup(id uint, userID uint) error {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, subgroup.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if !helpers.IsManagerOrModeratorRole(role) {
		return errors.New("insufficient permissions")
	}

	return u.subgroupRepo.DeleteSubgroup(id)
}

// AddMembers adds specified users to a subgroup.
func (u *SubgroupUsecase) AddMembers(id uint, userIDs []uint, requestingUserID uint) error {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(requestingUserID, subgroup.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if !helpers.IsManagerOrModeratorRole(role) {
		return errors.New("insufficient permissions")
	}

	return u.subgroupRepo.AddMembers(id, userIDs)
}

// RemoveMember removes a user from a subgroup.
func (u *SubgroupUsecase) RemoveMember(subgroupID uint, userID uint, requestingUserID uint) error {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(subgroupID)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(requestingUserID, subgroup.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != helpers.RoleManager {
		return errors.New("insufficient permissions")
	}

	return u.subgroupRepo.RemoveMember(subgroupID, userID)
}

// GetGroupSubgroups retrieves all subgroups in a specific group.
func (u *SubgroupUsecase) GetGroupSubgroups(groupID uint, userID uint) ([]*model.Subgroup, error) {
	_, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return u.subgroupRepo.GetGroupSubgroups(groupID)
}
