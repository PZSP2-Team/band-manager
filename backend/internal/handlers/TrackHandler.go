package handlers

import (
	"band-manager-backend/internal/usecases"
)

type TrackHandler struct {
	trackUsecase *usecases.TrackUsecase
}

func NewTrackHandler() *TrackHandler {
	return &TrackHandler{
		trackUsecase: usecases.NewTrackUsecase(),
	}
}
