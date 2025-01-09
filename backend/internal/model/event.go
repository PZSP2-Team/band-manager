package model

import "time"

type Event struct {
	ID           uint   `gorm:"primarykey"`
	Title        string `gorm:"not null"`
	Location     string `gorm:"not null"`
	Description  string
	Date         time.Time
	GroupID      uint          `gorm:"not null"`
	Group        Group         `gorm:"foreignKey:GroupID" json:"group"`
	Tracks       []*Track      `gorm:"many2many:event_tracks;" json:"tracks"`
	Performances []Performance `gorm:"constraint:OnDelete:CASCADE"`
}
