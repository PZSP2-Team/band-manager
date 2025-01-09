package model

type Group struct {
	ID            uint   `gorm:"primarykey"`
	Name          string `gorm:"not null"`
	AccessToken   string `gorm:"unique;not null"`
	Description   string
	Users         []*User         `gorm:"many2many:user_group;constraint:OnDelete:CASCADE"`
	Subgroups     []Subgroup      `gorm:"constraint:OnDelete:CASCADE"`
	Announcements []Announcement  `gorm:"constraint:OnDelete:CASCADE"`
	Events        []Event         `gorm:"constraint:OnDelete:CASCADE"`
	Tracks        []Track         `gorm:"constraint:OnDelete:CASCADE"`
	UserRoles     []UserGroupRole `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
}
