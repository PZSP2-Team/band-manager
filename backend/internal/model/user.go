package model

// User represents a system user.
type User struct {
	ID            uint            `gorm:"primarykey" json:"id"`
	FirstName     string          `gorm:"not null" json:"first_name"`
	LastName      string          `gorm:"not null" json:"last_name"`
	Email         string          `gorm:"unique;not null" json:"email"`
	PasswordHash  string          `gorm:"not null" json:"password_hash"`
	Groups        []*Group        `gorm:"many2many:user_group;constraint:OnDelete:CASCADE" json:"groups"`
	Announcements []Announcement  `gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL" json:"announcements"`
	Subgroups     []*Subgroup     `gorm:"many2many:subgroup_user;constraint:OnDelete:CASCADE" json:"subgroups"`
	GroupRoles    []UserGroupRole `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"group_roles"`
}
