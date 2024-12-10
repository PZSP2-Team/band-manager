package usecases

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"band-manager-backend/internal/models"
	"band-manager-backend/internal/repositories"
)

type UserUsecase interface {
	RegisterUser(name, email, password string) error
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) RegisterUser(name, email, password string) error {
	// Sprawdzenie, czy użytkownik już istnieje
	_, err := u.repo.FindByEmail(email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hashowanie hasła
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Tworzenie użytkownika
	user := &models.User{Name: name, Email: email, Password: string(hashedPassword)}
	return u.repo.Create(user)
}