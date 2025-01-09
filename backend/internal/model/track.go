package model

type Track struct {
	ID           uint          `gorm:"primarykey"`
	Name         string        `gorm:"not null"`
	GroupID      uint          `gorm:"not null"`
	Description  string        `gorm:"not null"`
	Events       []*Event      `gorm:"many2many:track_event;constraint:OnDelete:CASCADE"`
	Notesheets   []Notesheet   `gorm:"constraint:OnDelete:CASCADE"`
	Performances []Performance `gorm:"constraint:OnDelete:CASCADE"`
	Group        Group         `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
}
