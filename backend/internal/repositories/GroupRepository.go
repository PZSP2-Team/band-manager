package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"

	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{
		db: db.GetDB(),
	}
}

// CreateGroup creates a new group in the database
func (r *GroupRepository) CreateGroup(group *model.Group) error {
	result := r.db.Create(group)
	if result.Error != nil {
		return errors.New("failed to create group")
	}
	return nil
}

// GetGroupByAccessToken finds a group by its access token
func (r *GroupRepository) GetGroupByAccessToken(accessToken string) (*model.Group, error) {
	var group model.Group
	result := r.db.Where("access_token = ?", accessToken).First(&group)
	if result.Error != nil {
		return nil, errors.New("group not found")
	}
	return &group, nil
}
