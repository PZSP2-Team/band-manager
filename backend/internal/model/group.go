package model

type Group struct {
	ID            uint            `gorm:"primarykey" json:"id"`
	Name          string          `gorm:"not null" json:"name"`
	AccessToken   string          `gorm:"unique;not null" json:"access_token"`
	Description   string          `json:"description"`
	Users         []*User         `gorm:"many2many:user_group;constraint:OnDelete:CASCADE" json:"users"`
	Subgroups     []Subgroup      `gorm:"constraint:OnDelete:CASCADE" json:"subgroups"`
	Announcements []Announcement  `gorm:"constraint:OnDelete:CASCADE" json:"announcements"`
	Events        []Event         `gorm:"constraint:OnDelete:CASCADE" json:"events"`
	Tracks        []Track         `gorm:"constraint:OnDelete:CASCADE" json:"tracks"`
	UserRoles     []UserGroupRole `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"user_roles"`
}
