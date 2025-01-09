// usecases/event_usecase.go
package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"errors"
	"time"
)

type EventUsecase struct {
	eventRepo *repositories.EventRepository
	groupRepo *repositories.GroupRepository
	trackRepo *repositories.TrackRepository
}

func NewEventUsecase() *EventUsecase {
	return &EventUsecase{
		eventRepo: repositories.NewEventRepository(),
		groupRepo: repositories.NewGroupRepository(),
		trackRepo: repositories.NewTrackRepository(),
	}
}

func (u *EventUsecase) CreateEvent(title, description, location string, date time.Time, groupID uint, trackIDs []uint, userID uint) (*model.Event, error) {
	role, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("user not in group")
	}

	if role != "manager" && role != "moderator" {
		return nil, errors.New("insufficient permissions")
	}

	event := &model.Event{
		Title:       title,
		Description: description,
		Location:    location,
		Date:        date,
		GroupID:     groupID,
	}

	if err := u.eventRepo.CreateEvent(event); err != nil {
		return nil, err
	}

	if len(trackIDs) > 0 {
		for _, trackID := range trackIDs {
			track, err := u.trackRepo.GetTrackByID(trackID)
			if err != nil {
				return event, err
			}
			if track.GroupID != groupID {
				return event, errors.New("track does not belong to this group")
			}
		}

		err = u.eventRepo.AddTracksToEvent(event.ID, trackIDs)
		if err != nil {
			return event, err
		}
	}

	return event, nil
}

func (u *EventUsecase) GetEvent(eventID uint, userID uint) (*model.Event, error) {
	event, err := u.eventRepo.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	_, err = u.groupRepo.GetUserRole(userID, event.GroupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return event, nil
}

func (u *EventUsecase) UpdateEvent(id uint, title, description, location string, date time.Time, trackIDs []uint, userID uint) error {
	event, err := u.eventRepo.GetEventByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, event.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != "manager" && role != "moderator" {
		return errors.New("insufficient permissions")
	}

	event.Title = title
	event.Description = description
	event.Location = location
	event.Date = date

	if err := u.eventRepo.UpdateEvent(event); err != nil {
		return err
	}

	if trackIDs != nil {
		for _, trackID := range trackIDs {
			track, err := u.trackRepo.GetTrackByID(trackID)
			if err != nil {
				return err
			}
			if track.GroupID != event.GroupID {
				return errors.New("track does not belong to this group")
			}
		}

		return u.eventRepo.AddTracksToEvent(event.ID, trackIDs)
	}

	return nil
}

func (u *EventUsecase) DeleteEvent(id uint, userID uint) error {
	event, err := u.eventRepo.GetEventByID(id)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, event.GroupID)
	if err != nil {
		return errors.New("access denied")
	}

	if role != "manager" && role != "moderator" {
		return errors.New("insufficient permissions")
	}

	return u.eventRepo.DeleteEvent(id)
}

func (u *EventUsecase) GetGroupEvents(groupID uint, userID uint) ([]*model.Event, error) {
	_, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return u.eventRepo.GetGroupEvents(groupID)
}

func (u *EventUsecase) GetUserEvents(userID uint) ([]*model.Event, error) {
	return u.eventRepo.GetUserEvents(userID)
}