package handlers

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// AnnouncementHandler processes HTTP requests related to announcements.
type AnnouncementHandler struct {
	announcementUsecase *usecases.AnnouncementUsecase
}

func NewAnnouncementHandler() *AnnouncementHandler {
	return &AnnouncementHandler{
		announcementUsecase: usecases.NewAnnouncementUsecase(),
	}
}

// Create handles POST /api/announcement/create
// Creates a new announcement with specified details and recipients.
func (h *AnnouncementHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		Priority     uint   `json:"priority"`
		GroupID      uint   `json:"group_id"`
		SenderID     uint   `json:"sender_id"`
		RecipientIDs []uint `json:"recipient_ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	announcement, err := h.announcementUsecase.CreateAnnouncement(
		request.Title,
		request.Description,
		request.Priority,
		request.GroupID,
		request.SenderID,
		request.RecipientIDs,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(announcement)
}

// Delete handles DELETE /api/announcement/{id}/{userId}
// Deletes an existing announcement if user has proper permissions.
func (h *AnnouncementHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	announcementID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.announcementUsecase.DeleteAnnouncement(uint(announcementID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Announcement deleted successfully",
	})
}

// GetUserAnnouncements handles GET /api/announcement/user/{userId}
// Returns all announcements available to the specified user.
func (h *AnnouncementHandler) GetUserAnnouncements(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.ParseUint(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	announcements, err := h.announcementUsecase.GetUserAnnouncements(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]*model.Announcement{
		"announcements": announcements,
	})
}

// GetGroupAnnouncements handles GET /api/announcement/group/{groupId}/{userId}
// Returns all announcements for the specified group if user is a member.
func (h *AnnouncementHandler) GetGroupAnnouncements(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	groupID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	announcements, err := h.announcementUsecase.GetGroupAnnouncements(uint(groupID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]*model.Announcement{
		"announcements": announcements,
	})
}
