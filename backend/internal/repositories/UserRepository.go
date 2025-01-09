package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("użytkownik nie znaleziony")
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return errors.New("nie udało się utworzyć użytkownika")
	}

	return nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return errors.New("nie udało się zaktualizować użytkownika")
	}

	return nil
}

func (r *UserRepository) DeleteUser(userID uint) error {
	result := r.db.Delete(&model.User{}, userID)
	if result.Error != nil {
		return errors.New("nie udało się usunąć użytkownika")
	}

	return nil
}

func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User

	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("użytkownik nie znaleziony")
	}

	return &user, nil
}

func (r *UserRepository) GetUserGroupRoles(userID uint) ([]model.UserGroupRole, error) {
	var roles []model.UserGroupRole
	result := r.db.Where("user_id = ?", userID).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
