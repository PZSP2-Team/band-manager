package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// GroupRepository handles database operations for groups.
type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{
		db: db.GetDB(),
	}
}

// AddUserToGroup adds a user to a group with specified role.
func (r *GroupRepository) AddUserToGroup(userID uint, groupID uint, role string) error {
	return r.db.Create(&model.UserGroupRole{
		UserID:  userID,
		GroupID: groupID,
		Role:    role,
	}).Error
}

// GetUserRole retrieves a user's role within a group.
func (r *GroupRepository) GetUserRole(userID uint, groupID uint) (string, error) {
	var role model.UserGroupRole
	err := r.db.Where("user_id = ? AND group_id = ?", userID, groupID).First(&role).Error
	if err != nil {
		return "", err
	}
	return role.Role, nil
}

// CreateGroup persists a new group to the database.
func (r *GroupRepository) CreateGroup(group *model.Group) error {
	result := r.db.Create(group)
	if result.Error != nil {
		return errors.New("failed to create group")
	}
	return nil
}

// GetGroupByAccessToken retrieves a group using its access token.
func (r *GroupRepository) GetGroupByAccessToken(accessToken string) (*model.Group, error) {
	var group model.Group
	result := r.db.Where("access_token = ?", accessToken).First(&group)
	if result.Error != nil {
		return nil, errors.New("group not found")
	}
	return &group, nil
}

// GetGroupByID retrieves a group by its ID with related users.
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

// GetGroupMembers retrieves all members of a group.
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

// RemoveUserFromGroup removes a user from a group.
func (r *GroupRepository) RemoveUserFromGroup(userID uint, groupID uint) error {

	return r.db.Delete(&model.UserGroupRole{}, "user_id = ? AND group_id = ?", userID, groupID).Error
}

// UpdateUserRole updates a user's role within a group.
func (r *GroupRepository) UpdateUserRole(userID uint, groupID uint, newRole string) error {
	return r.db.Model(&model.UserGroupRole{}).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		Update("role", newRole).Error
}

// UpdateAccessToken updates a group's access token.
func (r *GroupRepository) UpdateAccessToken(groupID uint, newToken string) error {
	return r.db.Model(&model.Group{}).
		Where("id = ?", groupID).
		Update("access_token", newToken).Error
}
