package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo *repositories.UserRepository
}

func NewAuthUsecase() *AuthUsecase {
	userRepo := repositories.NewUserRepository()
	return &AuthUsecase{
		userRepo: userRepo,
	}
}

func (u *AuthUsecase) Login(email, password string) (*model.User, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}