package model

// Notesheet represents sheet music associated with a track.
type Notesheet struct {
	ID         uint        `gorm:"primarykey" json:"id"`
	Filepath   string      `gorm:"not null" json:"filepath"`
	Instrument string      `gorm:"not null" json:"instrument"`
	TrackId    uint        `gorm:"not null;" json:"track_id"`
	Subgroups  []*Subgroup `gorm:"many2many:notesheet_subgroup;constraint:OnDelete:CASCADE" json:"subgroups"`
	Track      Track       `gorm:"foreignKey:TrackId;constraint:OnDelete:CASCADE" json:"track"`
	FileType   string      `json:"file_type"`
	FileName   string      `json:"file_name"`
}
