package model

type Group struct {
	ID            uint   `gorm:"primarykey"`
	Name          string `gorm:"unique;not null"`
	AccessToken	  string `gorm:"unique;not null"`
	Description   string
	Users         []User         `gorm:"constraint:OnDelete:SET NULL"`
	Subgroups     []Subgroup     `gorm:"constraint:OnDelete:CASCADE"`
	Announcements []Announcement `gorm:"constraint:OnDelete:CASCADE"`
	Events        []Event        `gorm:"constraint:OnDelete:CASCADE"`
	Tracks        []Track        `gorm:"constraint:OnDelete:CASCADE"`
}
