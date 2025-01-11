package model

import "time"

type GoogleCalendarEvent struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	EventID    uint      `gorm:"not null;uniqueIndex" json:"event_id"`
	CalendarID string    `gorm:"not null" json:"calendar_id"`
	LastSynced time.Time `json:"last_synced"`
}
