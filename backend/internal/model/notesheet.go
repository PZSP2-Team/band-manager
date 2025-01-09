package model

type Notesheet struct {
	ID         uint        `gorm:"primarykey"`
	Filepath   string      `gorm:"not null"`
	Instrument string      `gorm:"not null"`
	TrackId    uint        `gorm:"not null;"`
	Subgroups  []*Subgroup `gorm:"many2many:notesheet_subgroup;constraint:OnDelete:CASCADE"`
}
