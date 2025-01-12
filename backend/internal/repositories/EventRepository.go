package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		db: db.GetDB(),
	}
}
func (r *EventRepository) GetEventUsers(eventID uint) ([]*model.User, error) {
	var event model.Event
	err := r.db.Preload("Users").First(&event, eventID).Error
	if err != nil {
		return nil, err
	}
	return event.Users, nil
}
func (r *EventRepository) AddUsersToEvent(eventID uint, userIDs []uint) error {
	var users []*model.User
	if err := r.db.Find(&users, userIDs).Error; err != nil {
		return err
	}
	return r.db.Model(&model.Event{ID: eventID}).Association("Users").Append(users)
}

func (r *EventRepository) CreateEvent(event *model.Event) error {
	return r.db.Create(event).Error
}

func (r *EventRepository) GetEventByID(id uint) (*model.Event, error) {
	var event model.Event
	if err := r.db.Preload("Group").Preload("Users").Preload("Tracks").Preload("Tracks.Notesheets").First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventRepository) UpdateEvent(event *model.Event) error {
	return r.db.Save(event).Error
}

func (r *EventRepository) DeleteEvent(id uint) error {
	return r.db.Delete(&model.Event{}, id).Error
}

func (r *EventRepository) GetGroupEvents(groupID uint) ([]*model.Event, error) {
	var events []*model.Event
	if err := r.db.Preload("Users").Where("group_id = ?", groupID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) GetUserEvents(userID uint) ([]*model.Event, error) {
	var events []*model.Event
	err := r.db.Preload("Users").
		Joins("JOIN event_users ON event_users.event_id = events.id").
		Where("event_users.user_id = ?", userID).
		Find(&events).Error
	return events, err
}

func (r *EventRepository) AddTracksToEvent(eventID uint, trackIDs []uint) error {
	var tracks []*model.Track
	if err := r.db.Find(&tracks, trackIDs).Error; err != nil {
		return err
	}

	return r.db.Model(&model.Event{ID: eventID}).Association("Tracks").Append(tracks)
}

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
