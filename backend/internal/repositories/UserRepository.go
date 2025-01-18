package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository handles database operations for users.
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

// GetUserByEmail retrieves a user by their email address.
func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("użytkownik nie znaleziony")
	}

	return &user, nil
}

// CreateUser persists a new user to the database.
func (r *UserRepository) CreateUser(user *model.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return errors.New("nie udało się utworzyć użytkownika")
	}

	return nil
}

// UpdateUser updates an existing user in the database.
func (r *UserRepository) UpdateUser(user *model.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return errors.New("nie udało się zaktualizować użytkownika")
	}

	return nil
}

// DeleteUser removes a user from the database.
func (r *UserRepository) DeleteUser(userID uint) error {
	result := r.db.Delete(&model.User{}, userID)
	if result.Error != nil {
		return errors.New("nie udało się usunąć użytkownika")
	}

	return nil
}

// GetUserByID retrieves a user by their ID.
func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User

	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("użytkownik nie znaleziony")
	}

	return &user, nil
}

// GetUserGroupRoles retrieves all group roles for a user.
func (r *UserRepository) GetUserGroupRoles(userID uint) ([]model.UserGroupRole, error) {
	var roles []model.UserGroupRole
	result := r.db.Where("user_id = ?", userID).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

// GetUserGroupRoles retrieves all group roles for a user.
func (r *UserRepository) ResetPassword(userID uint, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return r.db.Model(&model.User{}).Where("id = ?", userID).Update("password_hash", string(hashedPassword)).Error
}

// GetTotalUsers returns the total number of users in the system.
func (r *UserRepository) GetTotalUsers() (int64, error) {
	var count int64
	err := r.db.Model(&model.User{}).Count(&count).Error
	return count, err
}

func (r *GroupRepository) GetTotalGroups() (int64, error) {
	var count int64
	err := r.db.Model(&model.Group{}).Count(&count).Error
	return count, err
}
