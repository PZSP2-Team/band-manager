package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"
	"fmt"

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

func (r *GroupRepository) CreateGroup(group *model.Group) error {
	result := r.db.Create(group)
	if result.Error != nil {
		return errors.New("failed to create group")
	}
	return nil
}

func (r *GroupRepository) GetGroupByAccessToken(accessToken string) (*model.Group, error) {
	var group model.Group
	result := r.db.Where("access_token = ?", accessToken).First(&group)
	if result.Error != nil {
		return nil, errors.New("group not found")
	}
	return &group, nil
}

func (r *GroupRepository) GetGroupByID(id uint) (*model.Group, error) {
	var group model.Group

	result := r.db.First(&group, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("grupa o ID %d nie istnieje", id)
		}
		return nil, fmt.Errorf("błąd podczas pobierania grupy: %v", result.Error)
	}

	return &group, nil
}
