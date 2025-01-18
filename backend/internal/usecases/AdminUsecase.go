package usecases

import (
	"band-manager-backend/internal/domain"
	"band-manager-backend/internal/repositories"
)

// AdminUsecase implements administrative operations.
type AdminUsecase struct {
	userRepo  *repositories.UserRepository
	groupRepo *repositories.GroupRepository
}

func NewAdminUsecase() *AdminUsecase {
	return &AdminUsecase{
		userRepo:  repositories.NewUserRepository(),
		groupRepo: repositories.NewGroupRepository(),
	}
}

// ResetUserPassword changes a user's password to the provided new password.
func (u *AdminUsecase) ResetUserPassword(userID uint, newPassword string) error {
	return u.userRepo.ResetPassword(userID, newPassword)
}

// GetSystemStats retrieves system-wide statistics.
func (u *AdminUsecase) GetSystemStats() (domain.SystemStats, error) {
	totalUsers, err := u.userRepo.GetTotalUsers()
	if err != nil {
		return domain.SystemStats{}, err
	}

	totalGroups, err := u.groupRepo.GetTotalGroups()
	if err != nil {
		return domain.SystemStats{}, err
	}

	return domain.SystemStats{
		TotalUsers:  totalUsers,
		TotalGroups: totalGroups,
	}, nil
}
