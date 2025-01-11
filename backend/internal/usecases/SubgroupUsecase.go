package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"errors"
)

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

func (u *SubgroupUsecase) CreateSubgroup(name, description string, groupID uint, userID uint) (*model.Subgroup, error) {
	role, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}

	if role != "manager" && role != "moderator" {
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

func (u *SubgroupUsecase) UpdateSubgroup(id uint, name, description string, userID uint) error {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, subgroup.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != "manager" && role != "moderator" {
		return errors.New("insufficient permissions")
	}

	subgroup.Name = name
	subgroup.Description = description

	return u.subgroupRepo.UpdateSubgroup(subgroup)
}

func (u *SubgroupUsecase) DeleteSubgroup(id uint, userID uint) error {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, subgroup.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != "manager" && role != "moderator" {
		return errors.New("insufficient permissions")
	}

	return u.subgroupRepo.DeleteSubgroup(id)
}

func (u *SubgroupUsecase) AddMembers(id uint, userIDs []uint, requestingUserID uint) error {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(requestingUserID, subgroup.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != "manager" && role != "moderator" {
		return errors.New("insufficient permissions")
	}

	return u.subgroupRepo.AddMembers(id, userIDs)
}

func (u *SubgroupUsecase) RemoveMember(subgroupID uint, userID uint, requestingUserID uint) error {
	subgroup, err := u.subgroupRepo.GetSubgroupByID(subgroupID)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(requestingUserID, subgroup.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != "manager" {
		return errors.New("insufficient permissions")
	}

	return u.subgroupRepo.RemoveMember(subgroupID, userID)
}

func (u *SubgroupUsecase) GetGroupSubgroups(groupID uint, userID uint) ([]*model.Subgroup, error) {
	_, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return u.subgroupRepo.GetGroupSubgroups(groupID)
}
