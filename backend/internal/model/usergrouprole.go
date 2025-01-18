package model

// UserGroupRole represents a user's role within a group.
type UserGroupRole struct {
	UserID  uint   `gorm:"primarykey" json:"user_id"`
	GroupID uint   `gorm:"primarykey" json:"group_id"`
	Role    string `gorm:"not null" json:"role"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	Group   Group  `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"group"`
}
