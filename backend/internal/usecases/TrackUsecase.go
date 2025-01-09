package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"errors"
)

type TrackUsecase struct {
	trackRepo    *repositories.TrackRepository
	groupRepo    *repositories.GroupRepository
	subgroupRepo *repositories.SubgroupRepository
}

func NewTrackUsecase() *TrackUsecase {
	return &TrackUsecase{
		trackRepo:    repositories.NewTrackRepository(),
		groupRepo:    repositories.NewGroupRepository(),
		subgroupRepo: repositories.NewSubgroupRepository(),
	}
}

func (u *TrackUsecase) CreateTrack(title, description string, groupID uint, userID uint) (*model.Track, error) {
	// Sprawdź uprawnienia
	role, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}

	if role != "manager" && role != "moderator" {
		return nil, errors.New("insufficient permissions")
	}

	track := &model.Track{
		Name:        title,
		Description: description,
		GroupID:     groupID,
	}

	if err := u.trackRepo.CreateTrack(track); err != nil {
		return nil, err
	}

	return track, nil
}

func (u *TrackUsecase) GetTrack(id uint, userID uint) (*model.Track, error) {
	track, err := u.trackRepo.GetTrackByID(id)
	if err != nil {
		return nil, err
	}

	// Sprawdź czy użytkownik ma dostęp do grupy tego utworu
	_, err = u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return track, nil
}

func (u *TrackUsecase) UpdateTrack(id uint, title string, description string, userID uint) error {
	track, err := u.trackRepo.GetTrackByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != "manager" && role != "moderator" {
		return errors.New("insufficient permissions")
	}

	track.Name = title
	track.Description = description

	return u.trackRepo.UpdateTrack(track)
}

func (u *TrackUsecase) DeleteTrack(id uint, userID uint) error {
	track, err := u.trackRepo.GetTrackByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != "manager" && role != "moderator" {
		return errors.New("insufficient permissions")
	}

	return u.trackRepo.DeleteTrack(id)
}
