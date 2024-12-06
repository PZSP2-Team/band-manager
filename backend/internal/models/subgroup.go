package models

type Subgroup struct {
	ID            uint   `gorm:"primarykey"`
	GroupID       uint   `gorm:"not null;"`
	Name          string `gorm:"not null"`
	Description   string
	Users         []*User         `gorm:"many2many:subgroup_user;constraint:OnDelete:CASCADE"`
	Notesheets    []*Notesheet    `gorm:"many2many:notesheet_subgroup;constraint:OnDelete:CASCADE"`
	Announcements []*Announcement `gorm:"many2many:announcement_subgroup;constraint:OnDelete:CASCADE"`
}
