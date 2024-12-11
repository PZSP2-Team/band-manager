package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"
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

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDB(), // Używamy funkcji GetDB() z pakietu db do uzyskania instancji połączenia z bazą danych
	}
}
