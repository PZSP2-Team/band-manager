package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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

type AuthGroupInfo struct {
	ID   uint
	Name string
	Role string
}

func (u *AuthUsecase) Login(email, password string) (*model.User, []AuthGroupInfo, error) {
	// Najpierw sprawdzamy czy użytkownik istnieje
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, nil, errors.New("user not found")
	}

	// Sprawdzamy hasło
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, nil, errors.New("invalid credentials")
	}

	// Pobieramy informacje o grupach użytkownika i jego rolach w nich
	var userGroups []AuthGroupInfo

	// Pobieramy grupy i role użytkownika
	roles, err := u.userRepo.GetUserGroupRoles(user.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get user roles: %v", err)
	}

	for _, role := range roles {
		group, err := u.groupRepo.GetGroupByID(role.GroupID)
		if err != nil {
			continue // Pomijamy grupę jeśli wystąpił błąd
		}

		userGroups = append(userGroups, AuthGroupInfo{
			ID:   group.ID,
			Name: group.Name,
			Role: role.Role,
		})
	}

	return user, userGroups, nil
}

func (u *AuthUsecase) Register(firstName, lastName, email, password string) error {
	// Check if the email is already taken
	_, err := u.userRepo.GetUserByEmail(email)
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
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
