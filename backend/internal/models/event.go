package models

import "time"

type Event struct {
	ID           uint   `gorm:"primarykey"`
	Title        string `gorm:"not null"`
	Location     string `gorm:"not null"`
	Description  string
	Date         time.Time
	GroupID      uint     `gorm:"not null"`
	Performances []Performance `gorm:"constraint:OnDelete:CASCADE"`
}
