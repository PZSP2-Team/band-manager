package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// AuthUsecase implements authentication and user management logic.
type AuthUsecase struct {
	userRepo  *repositories.UserRepository
	groupRepo *repositories.GroupRepository
}

func NewAuthUsecase() *AuthUsecase {
	userRepo := repositories.NewUserRepository()
	groupRepo := repositories.NewGroupRepository()
	return &AuthUsecase{
		userRepo:  userRepo,
		groupRepo: groupRepo,
	}
}

// Login authenticates a user and returns their profile with group memberships.
func (u *AuthUsecase) Login(email, password string) (*model.User, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// Register creates a new user account with the provided details.
func (u *AuthUsecase) Register(firstName, lastName, email, password string) error {
	user, err := u.userRepo.GetUserByEmail(email)
	if user != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("incorrect password")
	}

	newUser := &model.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err := u.userRepo.CreateUser(newUser); err != nil {
		return errors.New("failed to create user")
	}

	return nil
}
