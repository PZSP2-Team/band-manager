package handlers

import (
	"band-manager-backend/internal/usecases"
)

type EventHandler struct {
	eventUsecase *usecases.EventUsecase
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		eventUsecase: usecases.NewEventUsecase(),
	}
}
