package usecases

import "band-manager-backend/internal/repositories"

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

func (u *AdminUsecase) ResetUserPassword(userID uint, newPassword string) error {
	return u.userRepo.ResetPassword(userID, newPassword)
}

type SystemStats struct {
	TotalUsers  int64 `json:"total_users"`
	TotalGroups int64 `json:"total_groups"`
}

func (u *AdminUsecase) GetSystemStats() (SystemStats, error) {
	totalUsers, err := u.userRepo.GetTotalUsers()
	if err != nil {
		return SystemStats{}, err
	}

	totalGroups, err := u.groupRepo.GetTotalGroups()
	if err != nil {
		return SystemStats{}, err
	}

	return SystemStats{
		TotalUsers:  totalUsers,
		TotalGroups: totalGroups,
	}, nil
}
