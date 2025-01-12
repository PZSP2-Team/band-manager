package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"errors"
)

// usecases/announcement_usecase.go
type AnnouncementUsecase struct {
	announcementRepo *repositories.AnnouncementRepository
	groupRepo        *repositories.GroupRepository
}

func NewAnnouncementUsecase() *AnnouncementUsecase {
	return &AnnouncementUsecase{
		announcementRepo: repositories.NewAnnouncementRepository(),
		groupRepo:        repositories.NewGroupRepository(),
	}
}

func (u *AnnouncementUsecase) CreateAnnouncement(title, description string, priority, groupID, senderID uint, recipientIDs []uint) (*model.Announcement, error) {
	role, err := u.groupRepo.GetUserRole(senderID, groupID)
	if err != nil {
		return nil, errors.New("could not get user role")
	}
	if role != "manager" && role != "moderator" {
		return nil, errors.New("insufficient permissions")
	}

	if len(recipientIDs) == 0 {
		groupUsers, err := u.groupRepo.GetGroupMembers(groupID)
		if err != nil {
			return nil, err
		}
		for _, user := range groupUsers {
			recipientIDs = append(recipientIDs, user.ID)
		}
	} else {
		for _, recipientID := range recipientIDs {
			_, err := u.groupRepo.GetUserRole(recipientID, groupID)
			if err != nil {
				return nil, errors.New("one or more recipients do not belong to the group")
			}
		}
	}

	announcement := &model.Announcement{
		Title:       title,
		Description: description,
		Priority:    priority,
		GroupID:     groupID,
		SenderID:    senderID,
	}

	if err := u.announcementRepo.Create(announcement); err != nil {
		return nil, err
	}

	if len(recipientIDs) > 0 {
		if err := u.announcementRepo.AddRecipients(announcement.ID, recipientIDs); err != nil {
			return nil, err
		}
	}

	return announcement, nil
}

func (u *AnnouncementUsecase) DeleteAnnouncement(announcementID, userID uint) error {
	announcement, err := u.announcementRepo.GetByID(announcementID)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, announcement.GroupID)
	if err != nil || (role != "manager" && role != "moderator" && announcement.SenderID != userID) {
		return errors.New("insufficient permissions")
	}

	return u.announcementRepo.Delete(announcementID)
}

func (u *AnnouncementUsecase) GetUserAnnouncements(userID uint) ([]*model.Announcement, error) {
	return u.announcementRepo.GetUserAnnouncements(userID)
}

func (u *AnnouncementUsecase) GetGroupAnnouncements(groupID, userID uint) ([]*model.Announcement, error) {
	_, err := u.groupRepo.GetUserRole(userID, groupID)
	if err != nil {
		return nil, errors.New("access denied")
	}
	return u.announcementRepo.GetGroupAnnouncements(groupID)
}
