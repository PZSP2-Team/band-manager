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

// EventHandler manages musical event operations.
type EventHandler struct {
	eventUsecase *usecases.EventUsecase
	gcService    *services.GoogleCalendarService
	emailService *services.EmailService
}

func NewEventHandler(gcService *services.GoogleCalendarService, emailService *services.EmailService) *EventHandler {
	return &EventHandler{
		eventUsecase: usecases.NewEventUsecase(gcService, emailService),
		gcService:    gcService,
		emailService: emailService,
	}
}

// Create handles POST /api/event/create
// Creates a new event with specified details, tracks, and participants.
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

// GetInfo handles GET /api/event/info/{eventId}/{userId}
// Retrieves detailed information about a specific event.
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

// Update handles PUT /api/event/update/{eventId}/{userId}
// Updates event details including tracks and participants.
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

// Delete handles DELETE /api/event/delete/{eventId}/{userId}
// Removes an event if user has proper permissions.
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

// GetGroupEvents handles GET /api/event/group/{groupId}/{userId}
// Returns all events for a specific group.
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

// GetUserEvents handles GET /api/event/user/{userId}
// Returns all events a user is participating in.
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

// GetEventTracks handles GET /api/event/tracks/{eventId}/{userId}
// Returns all tracks associated with an event.
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

// GoogleCalendarAuth handles GET /api/calendar/auth
// Initiates Google Calendar authentication flow.
func (h *EventHandler) GoogleCalendarAuth(w http.ResponseWriter, r *http.Request) {
	authURL := h.gcService.GetAuthURL()
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

// GoogleCalendarCallback handles GET /api/calendar/callback
// Processes Google Calendar authentication callback.
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
