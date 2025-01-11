package model

import "time"

type Performance struct {
	EventID   uint      `gorm:"primaryKey;not null" json:"event_id"`
	TrackID   uint      `gorm:"primaryKey;not null" json:"track_id"`
	StartTime time.Time `gorm:"not null" json:"start_time"`
}
