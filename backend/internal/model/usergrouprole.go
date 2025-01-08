package model

type UserGroupRole struct {
	UserID  uint   `gorm:"primarykey"`
	GroupID uint   `gorm:"primarykey"`
	Role    string `gorm:"not null"` // "manager", "moderator", "member"
}
