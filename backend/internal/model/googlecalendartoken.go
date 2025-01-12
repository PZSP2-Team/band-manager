package model

import "time"

type GoogleCalendarToken struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	UserID       uint      `gorm:"not null;uniqueIndex" json:"user_id"`
	AccessToken  string    `gorm:"not null" json:"access_token"`
	RefreshToken string    `gorm:"not null" json:"refresh_token"`
	ExpiryTime   time.Time `gorm:"not null" json:"expiry_time"`
	User         User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}
