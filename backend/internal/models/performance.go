package models

import "time"

type Performance struct {
    EventID   uint      `gorm:"primaryKey;not null"`
    TrackID   uint      `gorm:"primaryKey;not null"`
    StartTime time.Time `gorm:"not null"`
}
