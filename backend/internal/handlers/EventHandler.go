package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"net/http"
	"time"
)

type EventHandler struct {
	eventUsecase *usecases.EventUsecase
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		eventUsecase: usecases.NewEventUsecase(),
	}
}

func (h *EventHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Location    string    `json:"location"`
		Date        time.Time `json:"date"`
		GroupID     uint      `json:"group_id"`
		TrackIDs    []uint    `json:"track_ids"`
		UserID      uint      `json:"user_id"` // Tymczasowo, później z JWT
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	event, err := h.eventUsecase.CreateEvent(
		request.Title,
		request.Description,
		request.Location,
		request.Date,
		request.GroupID,
		request.TrackIDs,
		request.UserID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
