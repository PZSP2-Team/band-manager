package models

type User struct {
	ID				uint 			`gorm:"primarykey"`
	FirstName		string			`gorm:"not null"`
	LastName		string			`gorm:"not null"`
	Email			string 			`gorm:"unique;not null"`
	PasswordHash	string 			`gorm:"not null"`
	Role			string			`gorm:"not null"`
	GroupID			uint
	Announcements	[]Announcement	`gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL"`
	Subgroups		[]*Subgroup		`gorm:"many2many:subgroup_user;constraint:OnDelete:CASCADE"`
}
