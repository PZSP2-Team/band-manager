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