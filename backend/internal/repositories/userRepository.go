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

// GetUserByEmail pobiera użytkownika na podstawie adresu email
func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	// Używamy GORM do wyszukiwania użytkownika po emailu
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		// Jeśli nie znaleziono użytkownika, zwracamy błąd
		return nil, errors.New("user not found")
	}
	// Zwracamy użytkownika
	return &user, nil
}
