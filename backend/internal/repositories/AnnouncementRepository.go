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
	var subgroups []*model.Subgroup
	for _, id := range subgroupIDs {
		subgroups = append(subgroups, &model.Subgroup{ID: id})
	}
	return r.db.Model(&model.Announcement{ID: announcementID}).
		Association("Subgroups").
		Append(subgroups)
}

func (r *AnnouncementRepository) GetGroupAnnouncements(groupID uint) ([]*model.Announcement, error) {
	var announcements []*model.Announcement
	err := r.db.Where("group_id = ? AND id NOT IN (SELECT announcement_id FROM announcement_subgroup)",
		groupID).
		Preload("Sender").
		Order("priority desc").
		Find(&announcements).Error
	return announcements, err
}

func (r *AnnouncementRepository) GetUserAnnouncements(userID uint) ([]*model.Announcement, error) {
	var announcements []*model.Announcement
	err := r.db.Distinct().
		Where("announcements.id IN (SELECT announcement_id FROM announcement_subgroup asg "+
			"JOIN subgroup_user su ON asg.subgroup_id = su.subgroup_id "+
			"WHERE su.user_id = ?)", userID).
		Preload("Sender").
		Preload("Group").
		Preload("Subgroups").
		Order("priority desc").
		Find(&announcements).Error

	if err != nil {
		return nil, err
	}
	return announcements, nil
}
