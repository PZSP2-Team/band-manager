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
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (u *AuthUsecase) Register(firstName, lastName, email, password string) error {
	// Check if the email is already taken
	_, err := u.userRepo.GetUserByEmail(email)
	if err == nil {
		return errors.New("email already registered")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Create a new user
	newUser := &model.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         "user", // Default role
	}

	// Save the user to the database
	if err := u.userRepo.CreateUser(newUser); err != nil {
		return errors.New("failed to create user")
	}

	return nil
}
