package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{
		db: db.GetDB(),
	}
}

// CreateGroup creates a new group in the database
func (r *GroupRepository) CreateGroup(group *model.Group) error {
	result := r.db.Create(group)
	if result.Error != nil {
		return errors.New("failed to create group")
	}
	return nil
}

// GetGroupByAccessToken finds a group by its access token
func (r *GroupRepository) GetGroupByAccessToken(accessToken string) (*model.Group, error) {
	var group model.Group
	result := r.db.Where("access_token = ?", accessToken).First(&group)
	if result.Error != nil {
		return nil, errors.New("group not found")
	}
	return &group, nil
}

func (r *GroupRepository) GetGroupByID(id uint) (*model.Group, error) {
	var group model.Group

	// Używamy GORM do pobrania grupy z bazy danych
	result := r.db.First(&group, id)

	// Sprawdzamy, czy wystąpił błąd podczas pobierania
	if result.Error != nil {
		// Jeśli nie znaleziono grupy, zwracamy specyficzny błąd
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("grupa o ID %d nie istnieje", id)
		}
		// Dla innych błędów zwracamy ogólny komunikat błędu
		return nil, fmt.Errorf("błąd podczas pobierania grupy: %v", result.Error)
	}

	return &group, nil
}
