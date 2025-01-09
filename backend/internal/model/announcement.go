package model

type Announcement struct {
	ID          uint   `gorm:"primarykey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Priority    uint   `gorm:"not null"`
	GroupID     uint   `gorm:"not null"`
	SenderID    uint   `gorm:"not null"`

	Group     Group       `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
	Sender    User        `gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL"`
	Subgroups []*Subgroup `gorm:"many2many:announcement_subgroup;constraint:OnDelete:CASCADE"`
}
