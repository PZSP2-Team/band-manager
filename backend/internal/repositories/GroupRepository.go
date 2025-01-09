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

func (r *GroupRepository) AddUserToGroup(userID uint, groupID uint, role string) error {
	return r.db.Create(&model.UserGroupRole{
		UserID:  userID,
		GroupID: groupID,
		Role:    role,
	}).Error
}

func (r *GroupRepository) GetUserRole(userID uint, groupID uint) (string, error) {
	var role model.UserGroupRole
	err := r.db.Where("user_id = ? AND group_id = ?", userID, groupID).First(&role).Error
	if err != nil {
		return "", err
	}
	return role.Role, nil
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
	result := r.db.Preload("Users").First(&group, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("grupa o ID %d nie istnieje", id)
		}
		return nil, fmt.Errorf("błąd podczas pobierania grupy: %v", result.Error)
	}
	return &group, nil
}

func (r *GroupRepository) GetGroupMembers(groupID uint) ([]*model.User, error) {
	var roles []model.UserGroupRole
	err := r.db.Where("group_id = ?", groupID).Preload("User").Find(&roles).Error
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for _, role := range roles {
		users = append(users, &role.User)
	}
	return users, nil
}

func (r *GroupRepository) RemoveUserFromGroup(userID uint, groupID uint) error {

	return r.db.Delete(&model.UserGroupRole{}, "user_id = ? AND group_id = ?", userID, groupID).Error
}
