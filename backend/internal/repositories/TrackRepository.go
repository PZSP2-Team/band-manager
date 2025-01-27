package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"

	"gorm.io/gorm"
)

// TrackRepository handles database operations for tracks and notesheets.
type TrackRepository struct {
	db *gorm.DB
}

func NewTrackRepository() *TrackRepository {
	return &TrackRepository{
		db: db.GetDB(),
	}
}

// CreateTrack persists a new track to the database.
func (r *TrackRepository) CreateTrack(track *model.Track) error {
	return r.db.Create(track).Error
}

// GetTrackByID retrieves a track by its ID with related entities.
func (r *TrackRepository) GetTrackByID(id uint) (*model.Track, error) {
	var track model.Track
	if err := r.db.Preload("Group").Preload("Notesheets").First(&track, id).Error; err != nil {
		return nil, err
	}
	return &track, nil
}

// UpdateTrack updates an existing track in the database.
func (r *TrackRepository) UpdateTrack(track *model.Track) error {
	return r.db.Save(track).Error
}

// DeleteTrack removes a track and its associated resources.
func (r *TrackRepository) DeleteTrack(id uint) error {
	var track model.Track
	if err := r.db.First(&track, id).Error; err != nil {
		return err
	}

	if err := r.db.Model(&track).Association("Events").Clear(); err != nil {
		return err
	}

	return r.db.Delete(&track).Error
}

// GetGroupTracks retrieves all tracks for a specific group.
func (r *TrackRepository) GetGroupTracks(groupID uint) ([]*model.Track, error) {
	var tracks []*model.Track
	if err := r.db.Where("group_id = ?", groupID).
		Preload("Notesheets").
		Find(&tracks).Error; err != nil {
		return nil, err
	}
	return tracks, nil
}

// AddNotesheetToTrack creates a new notesheet and associates it with a track.
func (r *TrackRepository) AddNotesheetToTrack(notesheet *model.Notesheet, subgroupIDs []uint) error {
	if err := r.db.Create(notesheet).Error; err != nil {
		return err
	}

	if len(subgroupIDs) > 0 {
		var subgroups []model.Subgroup
		if err := r.db.Find(&subgroups, subgroupIDs).Error; err != nil {
			return err
		}
		return r.db.Model(notesheet).Association("Subgroups").Append(&subgroups)
	}
	return nil
}

// GetUserNotesheets retrieves notesheets available to a specific user.
func (r *TrackRepository) GetUserNotesheets(trackID, userID uint) ([]*model.Notesheet, error) {
	var notesheets []*model.Notesheet
	err := r.db.Joins("JOIN notesheet_subgroup ON notesheets.id = notesheet_subgroup.notesheet_id").
		Joins("JOIN subgroup_user ON notesheet_subgroup.subgroup_id = subgroup_user.subgroup_id").
		Where("subgroup_user.user_id = ? AND notesheets.track_id = ?", userID, trackID).
		Distinct().
		Find(&notesheets).Error
	return notesheets, err
}

// GetTrackNotesheets retrieves all notesheets for a track.
func (r *TrackRepository) GetTrackNotesheets(trackID uint) ([]*model.Notesheet, error) {
	var notesheets []*model.Notesheet
	err := r.db.Where("track_id = ?", trackID).Find(&notesheets).Error
	return notesheets, err
}

// UpdateNotesheetFilepath updates the file path of a notesheet.
func (r *TrackRepository) UpdateNotesheetFilepath(notesheetID uint, filepath string) error {
	return r.db.Model(&model.Notesheet{}).
		Where("id = ?", notesheetID).
		Update("filepath", filepath).
		Error
}

// GetNotesheet retrieves a notesheet by its ID.
func (r *TrackRepository) GetNotesheet(id uint) (*model.Notesheet, error) {
	var notesheet model.Notesheet
	if err := r.db.First(&notesheet, id).Error; err != nil {
		return nil, err
	}
	return &notesheet, nil
}
