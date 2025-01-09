// usecases/event_usecase.go
package usecases

import (
	"band-manager-backend/internal/repositories"
)

type EventUsecase struct {
	eventRepo *repositories.EventRepository
	groupRepo *repositories.GroupRepository
	trackRepo *repositories.TrackRepository
}

func NewEventUsecase() *EventUsecase {
	return &EventUsecase{
		eventRepo: repositories.NewEventRepository(),
		groupRepo: repositories.NewGroupRepository(),
		trackRepo: repositories.NewTrackRepository(),
	}
}
