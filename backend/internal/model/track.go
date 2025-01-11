package model

type Track struct {
	ID           uint          `gorm:"primarykey" json:"id"`
	Name         string        `gorm:"not null" json:"name"`
	GroupID      uint          `gorm:"not null" json:"group_id"`
	Description  string        `gorm:"not null" json:"description"`
	Events       []*Event      `gorm:"many2many:track_event;constraint:OnDelete:CASCADE" json:"events"`
	Notesheets   []Notesheet   `gorm:"constraint:OnDelete:CASCADE" json:"notesheets"`
	Performances []Performance `gorm:"constraint:OnDelete:CASCADE" json:"performances"`
	Group        Group         `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"group"`
}
