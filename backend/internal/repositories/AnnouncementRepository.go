package repositories

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"

	"gorm.io/gorm"
)

type AnnouncementRepository struct {
	db *gorm.DB
}

func NewAnnouncementRepository() *AnnouncementRepository {
	return &AnnouncementRepository{
		db: db.GetDB(),
	}
}

func (r *AnnouncementRepository) Create(announcement *model.Announcement) error {
	return r.db.Create(announcement).Error
}

func (r *AnnouncementRepository) GetByID(id uint) (*model.Announcement, error) {
	var announcement model.Announcement
	err := r.db.Preload("Group").Preload("Sender").Preload("Subgroups").First(&announcement, id).Error
	if err != nil {
		return nil, err
	}
	return &announcement, nil
}

func (r *AnnouncementRepository) Delete(id uint) error {
	return r.db.Delete(&model.Announcement{}, id).Error
}

func (r *AnnouncementRepository) AddToSubgroups(announcementID uint, subgroupIDs []uint) error {
	err := r.db.Model(&model.Announcement{ID: announcementID}).
		Association("Subgroups").
		Append(&model.Subgroup{ID: subgroupIDs[0]})
	return err
}
func (r *AnnouncementRepository) GetGroupAnnouncements(groupID uint) ([]model.Announcement, error) {
	var announcements []model.Announcement
	err := r.db.Where("group_id = ?", groupID).
		Preload("Sender").
		Order("priority desc, created_at desc").
		Find(&announcements).Error
	return announcements, err
}
