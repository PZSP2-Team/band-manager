package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"

	"gorm.io/gorm"
)

// UserRepository struktura przechowująca dostęp do bazy danych
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository tworzy nową instancję UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

// GetUserByEmail znajduje użytkownika po adresie email
func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("użytkownik nie znaleziony")
	}

	return &user, nil
}

// CreateUser tworzy nowego użytkownika w bazie danych
func (r *UserRepository) CreateUser(user *model.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return errors.New("nie udało się utworzyć użytkownika")
	}

	return nil
}

// UpdateUser aktualizuje dane użytkownika
func (r *UserRepository) UpdateUser(user *model.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return errors.New("nie udało się zaktualizować użytkownika")
	}

	return nil
}

// DeleteUser usuwa użytkownika z bazy danych
func (r *UserRepository) DeleteUser(userID uint) error {
	result := r.db.Delete(&model.User{}, userID)
	if result.Error != nil {
		return errors.New("nie udało się usunąć użytkownika")
	}

	return nil
}

// GetUserByID znajduje użytkownika po ID
func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User

	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("użytkownik nie znaleziony")
	}

	return &user, nil
}
