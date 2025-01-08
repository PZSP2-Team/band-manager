package model

type UserGroupRole struct {
	UserID  uint   `gorm:"primarykey"`
	GroupID uint   `gorm:"primarykey"`
	Role    string `gorm:"not null"` // "manager", "moderator", "member"

	// Dodanie relacji
	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Group Group `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
}
