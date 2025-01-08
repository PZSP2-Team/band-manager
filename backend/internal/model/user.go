package model

type User struct {
	ID           uint   `gorm:"primarykey"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	// Role         string
	// GroupID       *uint
	Groups        []*Group        `gorm:"many2many:user_group;constraint:OnDelete:CASCADE"`
	Announcements []Announcement  `gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL"`
	Subgroups     []*Subgroup     `gorm:"many2many:subgroup_user;constraint:OnDelete:CASCADE"`
	GroupRoles    []UserGroupRole `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
