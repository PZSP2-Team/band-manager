package model

import "time"

// Announcement represents a message sent to group or subgroup members.
type Announcement struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Priority    uint      `gorm:"not null" json:"priority"`
	GroupID     uint      `gorm:"not null" json:"group_id"`
	SenderID    uint      `gorm:"not null" json:"sender_id"`
	Group       Group     `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"group"`
	Sender      User      `gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL" json:"sender"`
	Recipients  []*User   `gorm:"many2many:announcement_recipients;constraint:OnDelete:CASCADE" json:"recipients"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
