package model

import "time"

type GoogleToken struct {
	ID           uint      `gorm:"primarykey"`
	UserID       uint      `gorm:"uniqueIndex"`
	AccessToken  string    `gorm:"not null"`
	TokenType    string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	Expiry       time.Time `gorm:"not null"`
}
