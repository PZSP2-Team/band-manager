package model

import "time"

type Announcement struct {
	ID          uint        `gorm:"primarykey" json:"id"`
	Title       string      `gorm:"not null" json:"title"`
	Description string      `gorm:"not null" json:"description"`
	Priority    uint        `gorm:"not null" json:"priority"`
	GroupID     uint        `gorm:"not null" json:"group_id"`
	SenderID    uint        `gorm:"not null" json:"sender_id"`
	Group       Group       `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"group"`
	Sender      User        `gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL" json:"sender"`
	Subgroups   []*Subgroup `gorm:"many2many:announcement_subgroup;constraint:OnDelete:CASCADE" json:"subgroups"`
	CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"` // Pole daty utworzenia
}
