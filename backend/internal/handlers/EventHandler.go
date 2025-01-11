package handlers

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/services"
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type EventHandler struct {
	eventUsecase *usecases.EventUsecase
	gcService    *services.GoogleCalendarService
}

func NewEventHandler(gcService *services.GoogleCalendarService) *EventHandler {
	return &EventHandler{
		eventUsecase: usecases.NewEventUsecase(gcService),
		gcService:    gcService,
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
		UserIDs     []uint    `json:"user_ids"`
		UserID      uint      `json:"user_id"`
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
		request.UserIDs,
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

func (h *EventHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	event, err := h.eventUsecase.GetEvent(uint(id), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var request struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Location    string    `json:"location"`
		Date        time.Time `json:"date"`
		TrackIDs    []uint    `json:"track_ids"`
		UserIDs     []uint    `json:"user_ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.eventUsecase.UpdateEvent(
		uint(id),
		request.Title,
		request.Description,
		request.Location,
		request.Date,
		request.TrackIDs,
		request.UserIDs,
		uint(userID),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Event updated successfully",
	})
}

func (h *EventHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.eventUsecase.DeleteEvent(uint(id), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Event deleted successfully",
	})
}

func (h *EventHandler) GetGroupEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	groupID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	events, err := h.eventUsecase.GetGroupEvents(uint(groupID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]*model.Event{
		"events": events,
	})
}

func (h *EventHandler) GetUserEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	events, err := h.eventUsecase.GetUserEvents(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]*model.Event{
		"events": events,
	})
}

func (h *EventHandler) GetEventTracks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	eventID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	tracks, err := h.eventUsecase.GetEventTracks(uint(eventID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]*model.Track{
		"tracks": tracks,
	})
}

func (h *EventHandler) GoogleCalendarAuth(w http.ResponseWriter, r *http.Request) {
	authURL := h.gcService.GetAuthURL()
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func (h *EventHandler) GoogleCalendarCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		return
	}

	fmt.Println("Authorization code received:", code) // DEBUG

	err := h.gcService.SaveToken(code)
	if err != nil {
		fmt.Println("Error saving token:", err) // DEBUG
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Token successfully saved!") // DEBUG
	w.Write([]byte("Successfully authorized!"))
}
