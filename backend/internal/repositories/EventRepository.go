package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"

	"gorm.io/gorm"
)

// EventRepository handles database operations for events.
type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		db: db.GetDB(),
	}
}

// GetEventUsers retrieves all users associated with an event.
func (r *EventRepository) GetEventUsers(eventID uint) ([]*model.User, error) {
	var event model.Event
	err := r.db.Preload("Users").First(&event, eventID).Error
	if err != nil {
		return nil, err
	}
	return event.Users, nil
}

// AddUsersToEvent associates users with an event.
func (r *EventRepository) AddUsersToEvent(eventID uint, userIDs []uint) error {
	var users []*model.User
	if err := r.db.Find(&users, userIDs).Error; err != nil {
		return err
	}
	return r.db.Model(&model.Event{ID: eventID}).Association("Users").Append(users)
}

// CreateEvent persists a new event to the database.
func (r *EventRepository) CreateEvent(event *model.Event) error {
	return r.db.Create(event).Error
}

// GetEventByID retrieves an event by its ID with related entities.
func (r *EventRepository) GetEventByID(id uint) (*model.Event, error) {
	var event model.Event
	if err := r.db.Preload("Group").Preload("Users").Preload("Tracks").First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

// UpdateEvent updates an existing event in the database.
func (r *EventRepository) UpdateEvent(event *model.Event) error {
	return r.db.Save(event).Error
}

// DeleteEvent removes an event from the database.
func (r *EventRepository) DeleteEvent(id uint) error {
	return r.db.Delete(&model.Event{}, id).Error
}

// GetGroupEvents retrieves all events for a specific group.
func (r *EventRepository) GetGroupEvents(groupID uint) ([]*model.Event, error) {
	var events []*model.Event
	if err := r.db.Preload("Users").Where("group_id = ?", groupID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

// GetGroupEvents retrieves all events for a specific group.
func (r *EventRepository) GetUserEvents(userID uint) ([]*model.Event, error) {
	var events []*model.Event
	err := r.db.Preload("Users").
		Joins("JOIN event_users ON event_users.event_id = events.id").
		Where("event_users.user_id = ?", userID).
		Find(&events).Error
	return events, err
}

// GetUserEvents retrieves all events a user is participating in.
func (r *EventRepository) AddTracksToEvent(eventID uint, trackIDs []uint) error {
	var tracks []*model.Track
	if err := r.db.Find(&tracks, trackIDs).Error; err != nil {
		return err
	}

	return r.db.Model(&model.Event{ID: eventID}).Association("Tracks").Append(tracks)
}

// AddTracksToEvent associates tracks with an event.
func (r *EventRepository) GetEventTracks(eventID uint) ([]*model.Track, error) {
	var event model.Event
	err := r.db.Preload("Tracks", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description")
	}).First(&event, eventID).Error
	if err != nil {
		return nil, err
	}
	return event.Tracks, nil
}
