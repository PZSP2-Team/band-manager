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

func (r *EventRepository) CreateEvent(event *model.Event) error {
	return r.db.Create(event).Error
}

func (r *EventRepository) GetEventByID(id uint) (*model.Event, error) {
	var event model.Event
	if err := r.db.Preload("Group").First(&event, id).Error; err != nil {
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
	if err := r.db.Where("group_id = ?", groupID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) GetUserEvents(userID uint) ([]*model.Event, error) {
	var events []*model.Event
	err := r.db.Joins("JOIN user_group_roles ON user_group_roles.group_id = events.group_id").
		Where("user_group_roles.user_id = ?", userID).
		Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) AddTracksToEvent(eventID uint, trackIDs []uint) error {
	var tracks []*model.Track
	if err := r.db.Find(&tracks, trackIDs).Error; err != nil {
		return err
	}

	return r.db.Model(&model.Event{ID: eventID}).Association("Tracks").Append(tracks)
}
