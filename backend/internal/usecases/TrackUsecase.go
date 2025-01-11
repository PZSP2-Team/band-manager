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

func (u *TrackUsecase) GetGroupTracks(groupID uint, userID uint) ([]*model.Track, error) {

	_, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return u.trackRepo.GetGroupTracks(groupID)
}

func (u *TrackUsecase) AddNotesheet(trackID uint, instrument string, filepath string, subgroupIDs []uint, userID uint) (*model.Notesheet, error) {
	track, err := u.trackRepo.GetTrackByID(trackID)
	if err != nil {
		return nil, err
	}

	role, err := u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}

	if role != "manager" && role != "moderator" {
		return nil, errors.New("insufficient permissions")
	}

	for _, subgroupID := range subgroupIDs {
		subgroup, err := u.subgroupRepo.GetSubgroupByID(subgroupID)
		if err != nil {
			return nil, err
		}
		if subgroup.GroupID != track.GroupID {
			return nil, errors.New("subgroup does not belong to track's group")
		}
	}

	notesheet := &model.Notesheet{
		TrackId:    trackID,
		Instrument: instrument,
		Filepath:   filepath,
	}

	if err := u.trackRepo.AddNotesheetToTrack(notesheet, subgroupIDs); err != nil {
		return nil, err
	}

	return notesheet, nil
}

func (u *TrackUsecase) GetUserNotesheets(trackID, userID uint) ([]*model.Notesheet, error) {
	return u.trackRepo.GetUserNotesheets(trackID, userID)
}

func (u *TrackUsecase) GetTrackNotesheets(trackID uint, userID uint) ([]*model.Notesheet, error) {
	track, err := u.trackRepo.GetTrackByID(trackID)
	if err != nil {
		return nil, err
	}

	// Sprawdź czy użytkownik ma dostęp do grupy
	_, err = u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return u.trackRepo.GetTrackNotesheets(trackID)
}

func (u *TrackUsecase) UpdateNotesheetFilepath(notesheetID uint, userID uint, filepath string) (*model.Notesheet, error) {
	// Pobierz notesheet
	notesheet, err := u.trackRepo.GetNotesheet(notesheetID)
	if err != nil {
		return nil, err
	}

	// Sprawdź uprawnienia
	track, err := u.trackRepo.GetTrackByID(notesheet.TrackId)
	if err != nil {
		return nil, err
	}

	role, err := u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}

	if role != "manager" && role != "moderator" {
		return nil, errors.New("insufficient permissions")
	}

	// Zaktualizuj ścieżkę do pliku
	err = u.trackRepo.UpdateNotesheetFilepath(notesheetID, filepath)
	if err != nil {
		return nil, err
	}

	return u.trackRepo.GetNotesheet(notesheetID)
}

func (u *TrackUsecase) GetNotesheet(notesheetID uint, userID uint) (*model.Notesheet, error) {
	// Pobierz notesheet
	notesheet, err := u.trackRepo.GetNotesheet(notesheetID)
	if err != nil {
		return nil, err
	}

	// Pobierz track, żeby sprawdzić uprawnienia
	track, err := u.trackRepo.GetTrackByID(notesheet.TrackId)
	if err != nil {
		return nil, err
	}

	// Sprawdź czy użytkownik ma dostęp do grupy
	_, err = u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return notesheet, nil
}
