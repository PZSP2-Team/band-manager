package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"
	"gorm.io/gorm"
)

// UserRepository definiuje interfejs dla operacji na użytkownikach
type UserRepository interface {
	// Pobiera użytkownika na podstawie emaila
	GetUserByEmail(email string) (*model.User, error)
	// Tworzy nowego użytkownika w bazie danych
	CreateUser(user *model.User) error
}

// userRepository to implementacja UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository tworzy nową instancję repozytorium użytkowników
func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetDB(), // Używamy funkcji GetDB() z pakietu db do uzyskania instancji połączenia z bazą danych
	}
}

// GetUserByEmail pobiera użytkownika na podstawie adresu email
func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	// Używamy GORM do wyszukiwania użytkownika po emailu
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		// Jeśli nie znaleziono użytkownika, zwracamy błąd
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		// W przypadku innych błędów, zwracamy ogólny błąd
		return nil, err
	}
	// Zwracamy użytkownika
	return &user, nil
}

// CreateUser tworzy nowego użytkownika w bazie danych
func (r *userRepository) CreateUser(user *model.User) error {
	// Używamy GORM do zapisu nowego użytkownika w bazie danych
	if err := r.db.Create(user).Error; err != nil {
		// Jeśli wystąpił błąd podczas zapisu, zwracamy błąd
		return err
	}
	// Zwracamy nil, gdy zapis był udany
	return nil
}
