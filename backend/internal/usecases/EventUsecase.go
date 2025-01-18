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

// EventUsecase implements event management logic.
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

// CreateEvent creates a new event with optional Google Calendar integration.
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

// DeleteEvent removes an event and its Google Calendar entry.
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

// GetGroupEvents retrieves all events for a specific group.
func (u *EventUsecase) GetGroupEvents(groupID uint, userID uint) ([]*model.Event, error) {
	if !u.isUserInGroup(userID, groupID) {
		return nil, errors.New("User not in group")
	}
	return u.eventRepo.GetGroupEvents(groupID)
}

// GetUserEvents retrieves all events a user is participating in.
func (u *EventUsecase) GetUserEvents(userID uint) ([]*model.Event, error) {
	return u.eventRepo.GetUserEvents(userID)
}

// GetEventTracks retrieves tracks associated with an event.
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

// GetEvent retrieves event details if user has access.
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

// UpdateEvent modifies event details and updates Google Calendar.
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

// Validates if the user has sufficient permissions (manager or moderator) in the group.
func (u *EventUsecase) validateUserPermissions(userID, groupID uint) error {
	role, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return errors.New("User not in group")
	}
	if !helpers.IsManagerOrModeratorRole(role) {
		return errors.New("insufficient permissions")
	}
	return nil
}

// Checks if the user is a member of the group.
func (u *EventUsecase) isUserInGroup(userID, groupID uint) bool {
	_, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return false
	}
	return true
}

// Adds tracks to the event, ensuring they belong to the same group.
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

// Adds users to the event, checking if they belong to the group.
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

// Handles external integrations like Google Calendar and email notifications.
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

// Returns a list of users to receive event notifications (either specific users or all group members).
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

// Updates basic event details (title, description, location, date).
func (u *EventUsecase) updateEventBasicInfo(event *model.Event, title, description, location string, date time.Time) error {
	event.Title = title
	event.Description = description
	event.Location = location
	event.Date = date

	return u.eventRepo.UpdateEvent(event)
}
