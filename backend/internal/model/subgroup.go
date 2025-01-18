package model

// Subgroup represents a subset of group members.
type Subgroup struct {
	ID            uint            `gorm:"primarykey" json:"id"`
	GroupID       uint            `gorm:"not null;" json:"group_id"`
	Name          string          `gorm:"not null" json:"name"`
	Description   string          `json:"description"`
	Users         []*User         `gorm:"many2many:subgroup_user;constraint:OnDelete:CASCADE" json:"users"`
	Notesheets    []*Notesheet    `gorm:"many2many:notesheet_subgroup;constraint:OnDelete:CASCADE" json:"notesheets"`
	Announcements []*Announcement `gorm:"many2many:announcement_subgroup;constraint:OnDelete:CASCADE" json:"announcements"`
}
