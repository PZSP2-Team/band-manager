package usecases

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"band-manager-backend/internal/services"
	"band-manager-backend/internal/usecases/helpers"
	"errors"
)

type AnnouncementUsecase struct {
	announcementRepo *repositories.AnnouncementRepository
	groupRepo        *repositories.GroupRepository
	emailService     *services.EmailService
	userRepo         *repositories.UserRepository
}

func NewAnnouncementUsecase() *AnnouncementUsecase {
	return &AnnouncementUsecase{
		announcementRepo: repositories.NewAnnouncementRepository(),
		groupRepo:        repositories.NewGroupRepository(),
		emailService:     services.NewEmailService(),
		userRepo:         repositories.NewUserRepository(),
	}
}

func (u *AnnouncementUsecase) CreateAnnouncement(title, description string, priority, groupID, senderID uint, recipientIDs []uint) (*model.Announcement, error) {
	role, err := u.groupRepo.GetUserRole(senderID, groupID)
	if err != nil {
		return nil, errors.New("Could not get user role")
	}
	if !helpers.IsManagerOrModeratorRole(role) {
		return nil, errors.New("insufficient permissions")
	}

	var recipients []*model.User
	if len(recipientIDs) == 0 {

		groupUsers, err := u.groupRepo.GetGroupMembers(groupID)
		if err != nil {
			return nil, err
		}
		recipients = groupUsers
		for _, user := range groupUsers {
			recipientIDs = append(recipientIDs, user.ID)
		}
	} else {

		for _, recipientID := range recipientIDs {

			_, err := u.groupRepo.GetUserRole(recipientID, groupID)
			if err != nil {
				return nil, errors.New("one or more recipients do not belong to the group")
			}

			user, err := u.userRepo.GetUserByID(recipientID)
			if err != nil {
				return nil, err
			}
			recipients = append(recipients, user)
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

	if err := u.announcementRepo.AddRecipients(announcement.ID, recipientIDs); err != nil {
		return nil, err
	}

	go u.emailService.SendAnnouncementEmail(announcement, recipients)

	return announcement, nil
}

func (u *AnnouncementUsecase) DeleteAnnouncement(announcementID, userID uint) error {
	announcement, err := u.announcementRepo.GetByID(announcementID)
	if err != nil {
		return err
	}

	role, err := u.groupRepo.GetUserRole(userID, announcement.GroupID)
	if err != nil {
		return err
	}
	if !helpers.IsManagerOrModeratorRole(role) && announcement.SenderID != userID {
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
		return nil, err
	}
	return u.announcementRepo.GetGroupAnnouncements(groupID)
}
