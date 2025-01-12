package repositories

import (
	"band-manager-backend/internal/infrastructure/db"
	"band-manager-backend/internal/model"

	"gorm.io/gorm"
)

type SubgroupRepository struct {
	db *gorm.DB
}

func NewSubgroupRepository() *SubgroupRepository {
	return &SubgroupRepository{
		db: db.GetDB(),
	}
}

func (r *SubgroupRepository) CreateSubgroup(subgroup *model.Subgroup) error {
	return r.db.Create(subgroup).Error
}

func (r *SubgroupRepository) GetSubgroupByID(id uint) (*model.Subgroup, error) {
	var subgroup model.Subgroup
	if err := r.db.Preload("Users").First(&subgroup, id).Error; err != nil {
		return nil, err
	}
	return &subgroup, nil
}

func (r *SubgroupRepository) UpdateSubgroup(subgroup *model.Subgroup) error {
	return r.db.Save(subgroup).Error
}

func (r *SubgroupRepository) DeleteSubgroup(id uint) error {
	return r.db.Delete(&model.Subgroup{}, id).Error
}

func (r *SubgroupRepository) AddMembers(subgroupID uint, userIDs []uint) error {
	var subgroup model.Subgroup
	if err := r.db.First(&subgroup, subgroupID).Error; err != nil {
		return err
	}

	var users []*model.User
	if err := r.db.Find(&users, userIDs).Error; err != nil {
		return err
	}

	return r.db.Model(&subgroup).Association("Users").Append(users)
}

func (r *SubgroupRepository) RemoveMember(subgroupID uint, userID uint) error {
	var subgroup model.Subgroup
	if err := r.db.First(&subgroup, subgroupID).Error; err != nil {
		return err
	}

	var user model.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	return r.db.Model(&subgroup).Association("Users").Delete(&user)
}

func (r *SubgroupRepository) GetGroupSubgroups(groupID uint) ([]*model.Subgroup, error) {
	var subgroups []*model.Subgroup
	if err := r.db.Preload("Users").Where("group_id = ?", groupID).Find(&subgroups).Error; err != nil {
		return nil, err
	}
	return subgroups, nil
}
