package model

import "time"

// Event represents a musical event or rehearsal.
type Event struct {
	ID                  uint                `gorm:"primarykey" json:"id"`
	Title               string              `gorm:"not null" json:"title"`
	Location            string              `gorm:"not null" json:"location"`
	Description         string              `json:"description"`
	Date                time.Time           `json:"date"`
	GroupID             uint                `gorm:"not null" json:"group_id"`
	Group               Group               `gorm:"foreignKey:GroupID" json:"group"`
	Tracks              []*Track            `gorm:"many2many:event_tracks;constraint:OnDelete:CASCADE" json:"tracks"`
	Users               []*User             `gorm:"many2many:event_users;constraint:OnDelete:CASCADE" json:"users"`
	Performances        []Performance       `gorm:"constraint:OnDelete:CASCADE" json:"performances"`
	GoogleCalendarEvent GoogleCalendarEvent `gorm:"constraint:OnDelete:CASCADE" json:"google_calendar_event,omitempty"`
}
