package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"band-manager-backend/internal/services"
	"band-manager-backend/internal/usecases/helpers"
	"errors"
	"log"
	"time"
)

type EventUsecase struct {
	eventRepo    *repositories.EventRepository
	groupRepo    *repositories.GroupRepository
	trackRepo    *repositories.TrackRepository
	userRepo     *repositories.UserRepository
	gcService    *services.GoogleCalendarService
	emailService *services.EmailService
}

func NewEventUsecase(gcService *services.GoogleCalendarService, emailService *services.EmailService) *EventUsecase {
	return &EventUsecase{
		eventRepo:    repositories.NewEventRepository(),
		groupRepo:    repositories.NewGroupRepository(),
		trackRepo:    repositories.NewTrackRepository(),
		userRepo:     repositories.NewUserRepository(),
		gcService:    gcService,
		emailService: emailService,
	}
}
func (u *EventUsecase) CreateEvent(title, description, location string, date time.Time,
	groupID uint, trackIDs []uint, userIDs []uint, userID uint) (*model.Event, error) {

	if err := u.validateUserPermissions(userID, groupID); err != nil {
		return nil, err
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

	if err := u.addTracksToEvent(event, trackIDs); err != nil {
		return event, err
	}

	if err := u.addUsersToEvent(event, userIDs); err != nil {
		return event, err
	}

	u.handleExternalIntegrations(event, userIDs)

	return event, nil
}

func (u *EventUsecase) DeleteEvent(id uint, userID uint) error {
	event, err := u.eventRepo.GetEventByID(id)
	if err != nil {
		return errors.New("Could not find event")
	}
	if err := u.validateUserPermissions(userID, event.GroupID); err != nil {
		return err
	}
	return u.eventRepo.DeleteEvent(id)
}

func (u *EventUsecase) GetGroupEvents(groupID uint, userID uint) ([]*model.Event, error) {
	if !u.isUserInGroup(userID, groupID) {
		return nil, errors.New("User not in group")
	}
	return u.eventRepo.GetGroupEvents(groupID)
}

func (u *EventUsecase) GetUserEvents(userID uint) ([]*model.Event, error) {
	return u.eventRepo.GetUserEvents(userID)
}

func (u *EventUsecase) GetEventTracks(eventID uint, userID uint) ([]*model.Track, error) {
	event, err := u.eventRepo.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	if !u.isUserInGroup(userID, event.GroupID) {
		return nil, errors.New("User not in group")
	}

	return u.eventRepo.GetEventTracks(eventID)
}

func (u *EventUsecase) GetEvent(eventID uint, userID uint) (*model.Event, error) {
	event, err := u.eventRepo.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	if !u.isUserInGroup(userID, event.GroupID) {
		return nil, errors.New("User not in group")
	}

	return event, nil
}

func (u *EventUsecase) UpdateEvent(id uint, title, description, location string,
	date time.Time, trackIDs []uint, userIDs []uint, userID uint) error {

	event, err := u.eventRepo.GetEventByID(id)
	if err != nil {
		return err
	}

	if err := u.validateUserPermissions(userID, event.GroupID); err != nil {
		return err
	}

	if err := u.updateEventBasicInfo(event, title, description, location, date); err != nil {
		return err
	}

	if trackIDs != nil {
		if err := u.addTracksToEvent(event, trackIDs); err != nil {
			return err
		}
	}

	if userIDs != nil {
		if err := u.addUsersToEvent(event, userIDs); err != nil {
			return err
		}
	}

	return nil
}
func (u *EventUsecase) validateUserPermissions(userID, groupID uint) error {
	if !u.isUserInGroup(userID, groupID) {
		return errors.New("User not in group")
	}
	if !helpers.HasManagerOrModeratorRole(role) {
		return errors.New("insufficient permissions")
	}
	return nil
}

func (u *EventUsecase) isUserInGroup(userID, groupID uint) bool {
	_, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return false
	}
	return true
}

func (u *EventUsecase) addTracksToEvent(event *model.Event, trackIDs []uint) error {
	if len(trackIDs) == 0 {
		return nil
	}

	for _, trackID := range trackIDs {
		track, err := u.trackRepo.GetTrackByID(trackID)
		if err != nil {
			return errors.New("Could not get track by ID")
		}
		if track.GroupID != event.GroupID {
			return errors.New("track does not belong to this group")
		}
	}

	return u.eventRepo.AddTracksToEvent(event.ID, trackIDs)
}

func (u *EventUsecase) addUsersToEvent(event *model.Event, userIDs []uint) error {
	if len(userIDs) > 0 {
		for _, assignedUserID := range userIDs {
			if !u.isUserInGroup(assignedUserID, event.GroupID) {
				return errors.New("Some users not in group")
			}
		}
		return u.eventRepo.AddUsersToEvent(event.ID, userIDs)
	}

	groupUsers, err := u.groupRepo.GetGroupMembers(event.GroupID)
	if err != nil {
		return err
	}
	var allUserIDs []uint
	for _, user := range groupUsers {
		allUserIDs = append(allUserIDs, user.ID)
	}
	return u.eventRepo.AddUsersToEvent(event.ID, allUserIDs)
}

func (u *EventUsecase) handleExternalIntegrations(event *model.Event, userIDs []uint) {
	if u.gcService != nil {
		if err := u.gcService.CreateCalendarEvent(event); err != nil {
			log.Printf("Failed to sync with Google Calendar: %v", err)
		}
	}

	recipients, err := u.getEventRecipients(event.GroupID, userIDs)
	if err != nil {
		log.Printf("Failed to get recipients: %v", err)
		return
	}

	if len(recipients) > 0 {
		go u.emailService.SendEventEmail(event, recipients)
	}
}

func (u *EventUsecase) getEventRecipients(groupID uint, userIDs []uint) ([]*model.User, error) {
	if len(userIDs) > 0 {
		var recipients []*model.User
		for _, id := range userIDs {
			user, err := u.userRepo.GetUserByID(id)
			if err != nil {
				continue
			}
			recipients = append(recipients, user)
		}
		return recipients, nil
	}

	return u.groupRepo.GetGroupMembers(groupID)
}

func (u *EventUsecase) updateEventBasicInfo(event *model.Event, title, description, location string, date time.Time) error {
	event.Title = title
	event.Description = description
	event.Location = location
	event.Date = date

	return u.eventRepo.UpdateEvent(event)
}
