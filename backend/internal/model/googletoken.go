package model

import "time"

// GoogleToken stores Google OAuth tokens for calendar integration.
type GoogleToken struct {
	ID           uint      `gorm:"primarykey"`
	UserID       uint      `gorm:"uniqueIndex"`
	AccessToken  string    `gorm:"not null"`
	TokenType    string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	Expiry       time.Time `gorm:"not null"`
}
