package models

type Announcement struct {
	ID          uint   `gorm:"primarykey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Priority    uint   `gorm:"not null"`
	GroupID     uint
	SenderID    uint
}
