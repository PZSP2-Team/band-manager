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

