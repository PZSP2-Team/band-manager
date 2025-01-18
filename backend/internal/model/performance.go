package model

import "time"

// Performance represents a track performed at an event.
type Performance struct {
	EventID   uint      `gorm:"primaryKey;not null" json:"event_id"`
	TrackID   uint      `gorm:"primaryKey;not null" json:"track_id"`
	StartTime time.Time `gorm:"not null" json:"start_time"`
}
