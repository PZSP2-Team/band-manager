package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"band-manager-backend/internal/usecases/helpers"
	"errors"
)

// TrackUsecase implements music track management logic.
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

// CreateTrack creates a new track in a specified group.
func (u *TrackUsecase) CreateTrack(title, description string, groupID uint, userID uint) (*model.Track, error) {
	role, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}
	if !helpers.IsManagerOrModeratorRole(role) {
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

// GetTrack retrieves track details if user has access.
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

// UpdateTrack modifies track details if user has permissions.
func (u *TrackUsecase) UpdateTrack(id uint, title string, description string, userID uint) error {
	track, err := u.trackRepo.GetTrackByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if !helpers.IsManagerOrModeratorRole(role) {
		return errors.New("insufficient permissions")
	}

	track.Name = title
	track.Description = description

	return u.trackRepo.UpdateTrack(track)
}

// DeleteTrack removes a track and associated resources.
func (u *TrackUsecase) DeleteTrack(id uint, userID uint) error {
	track, err := u.trackRepo.GetTrackByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if !helpers.IsManagerOrModeratorRole(role) {
		return errors.New("insufficient permissions")
	}

	return u.trackRepo.DeleteTrack(id)
}

// GetGroupTracks retrieves all tracks in a specific group.
func (u *TrackUsecase) GetGroupTracks(groupID uint, userID uint) ([]*model.Track, error) {

	_, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return u.trackRepo.GetGroupTracks(groupID)
}

// AddNotesheet adds a new notesheet to a track for specific subgroups.
func (u *TrackUsecase) AddNotesheet(trackID uint, instrument string, filepath string, subgroupIDs []uint, userID uint) (*model.Notesheet, error) {
	track, err := u.trackRepo.GetTrackByID(trackID)
	if err != nil {
		return nil, err
	}

	role, err := u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}
	if !helpers.IsManagerOrModeratorRole(role) {
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

// GetUserNotesheets retrieves notesheets available to a specific user.
func (u *TrackUsecase) GetUserNotesheets(trackID, userID uint) ([]*model.Notesheet, error) {
	return u.trackRepo.GetUserNotesheets(trackID, userID)
}

// GetTrackNotesheets retrieves all notesheets for a track.
func (u *TrackUsecase) GetTrackNotesheets(trackID uint, userID uint) ([]*model.Notesheet, error) {
	track, err := u.trackRepo.GetTrackByID(trackID)
	if err != nil {
		return nil, err
	}

	_, err = u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return u.trackRepo.GetTrackNotesheets(trackID)
}

// UpdateNotesheetFilepath updates the file path for a notesheet.
func (u *TrackUsecase) UpdateNotesheetFilepath(notesheetID uint, userID uint, filepath string) (*model.Notesheet, error) {
	notesheet, err := u.trackRepo.GetNotesheet(notesheetID)
	if err != nil {
		return nil, err
	}

	track, err := u.trackRepo.GetTrackByID(notesheet.TrackId)
	if err != nil {
		return nil, err
	}

	role, err := u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}

	if !helpers.IsManagerOrModeratorRole(role) {
		return nil, errors.New("insufficient permissions")
	}

	err = u.trackRepo.UpdateNotesheetFilepath(notesheetID, filepath)
	if err != nil {
		return nil, err
	}

	return u.trackRepo.GetNotesheet(notesheetID)
}

// GetNotesheet retrieves notesheet details if user has access.
func (u *TrackUsecase) GetNotesheet(notesheetID uint, userID uint) (*model.Notesheet, error) {
	notesheet, err := u.trackRepo.GetNotesheet(notesheetID)
	if err != nil {
		return nil, err
	}

	track, err := u.trackRepo.GetTrackByID(notesheet.TrackId)
	if err != nil {
		return nil, err
	}

	_, err = u.groupRepo.GetUserRole(userID, track.GroupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return notesheet, nil
}
