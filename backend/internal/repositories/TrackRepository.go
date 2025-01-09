package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"

	"gorm.io/gorm"
)

type TrackRepository struct {
	db *gorm.DB
}

func NewTrackRepository() *TrackRepository {
	return &TrackRepository{
		db: db.GetDB(),
	}
}

func (r *TrackRepository) CreateTrack(track *model.Track) error {
	return r.db.Create(track).Error
}

func (r *TrackRepository) GetTrackByID(id uint) (*model.Track, error) {
	var track model.Track
	if err := r.db.Preload("Group").Preload("Notesheets").First(&track, id).Error; err != nil {
		return nil, err
	}
	return &track, nil
}

func (r *TrackRepository) UpdateTrack(track *model.Track) error {
	return r.db.Save(track).Error
}

func (r *TrackRepository) DeleteTrack(id uint) error {
	return r.db.Delete(&model.Track{}, id).Error
}

func (r *TrackRepository) GetGroupTracks(groupID uint) ([]*model.Track, error) {
	var tracks []*model.Track
	if err := r.db.Where("group_id = ?", groupID).
		Preload("Notesheets").
		Find(&tracks).Error; err != nil {
		return nil, err
	}
	return tracks, nil
}

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

func (r *TrackRepository) GetSubgroupNotesheets(subgroupID uint) ([]*model.Notesheet, error) {
	var notesheets []*model.Notesheet
	err := r.db.Joins("JOIN notesheet_subgroup ON notesheets.id = notesheet_subgroup.notesheet_id").
		Where("notesheet_subgroup.subgroup_id = ?", subgroupID).
		Find(&notesheets).Error
	return notesheets, err
}
