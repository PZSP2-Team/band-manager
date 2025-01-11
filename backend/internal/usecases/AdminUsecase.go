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

func (u *AdminUsecase) GetSystemStats() (map[string]interface{}, error) {
	totalUsers, err := u.userRepo.GetTotalUsers()
	if err != nil {
		return nil, err
	}

	totalGroups, err := u.groupRepo.GetTotalGroups()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_users":  totalUsers,
		"total_groups": totalGroups,
	}, nil
}
