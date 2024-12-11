package repositories

import (
	"band-manager-backend/internal/models"

	"gorm.io/gorm"
)

// UserRepository definiuje interfejs dla operacji na użytkownikach
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

// userRepository to implementacja UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository tworzy nowy obiekt UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}