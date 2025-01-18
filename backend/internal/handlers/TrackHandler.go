package handlers

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const UPLOAD_DIR = "/app/uploads"

// TrackHandler manages musical track operations.
type TrackHandler struct {
	trackUsecase *usecases.TrackUsecase
}

func NewTrackHandler() *TrackHandler {
	return &TrackHandler{
		trackUsecase: usecases.NewTrackUsecase(),
	}
}

// Create handles POST /api/track/create
// Creates a new track in the specified group.
func (h *TrackHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		GroupID     uint   `json:"group_id"`
		UserID      uint   `json:"user_id"` // Tymczasowo, później z JWT
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	track, err := h.trackUsecase.CreateTrack(
		request.Title,
		request.Description,
		request.GroupID,
		request.UserID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(track)
}

// AddNotesheet handles POST /api/track/notesheet
// Adds a new notesheet to a track for specific subgroups.
func (h *TrackHandler) AddNotesheet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		TrackID     uint   `json:"track_id"`
		UserID      uint   `json:"user_id"` // typ MIME pliku
		FileName    string // oryginalna nazwa pliku
		Filepath    string `json:"filepath"`
		Instrument  string `json:"instrument"`
		SubgroupIDs []uint `json:"subgroup_ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	notesheet, err := h.trackUsecase.AddNotesheet(
		request.TrackID,
		"default",
		request.Filepath,
		request.SubgroupIDs,
		request.UserID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(notesheet)
}

// GetUserNotesheets handles GET /api/track/user/notesheets/{trackId}/{userId}
// Returns notesheets available to a specific user.
func (h *TrackHandler) GetUserNotesheets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	trackID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	notesheets, err := h.trackUsecase.GetUserNotesheets(uint(trackID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"notesheets": notesheets,
	})
}

// GetGroupTracks handles GET /api/track/group/{groupId}/{userId}
// Returns all tracks in a group.
func (h *TrackHandler) GetGroupTracks(w http.ResponseWriter, r *http.Request) {
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

	tracks, err := h.trackUsecase.GetGroupTracks(uint(groupID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tracks": tracks,
	})
}

// GetTrackNotesheets handles GET /api/track/notesheets/{trackId}/{userId}
// Returns all notesheets for a track.
func (h *TrackHandler) GetTrackNotesheets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	trackID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	notesheets, err := h.trackUsecase.GetTrackNotesheets(uint(trackID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]*model.Notesheet{
		"notesheets": notesheets,
	})
}

// UploadNotesheetFile handles POST /api/track/notesheet/upload/{notesheetId}/{userId}
// Handles file upload for a notesheet.
func (h *TrackHandler) UploadNotesheetFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	notesheetID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid notesheet ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(10 << 20) // 10MB limit

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), handler.Filename)
	filepath := path.Join(UPLOAD_DIR, filename)

	if err := os.MkdirAll(UPLOAD_DIR, 0755); err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(filepath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	notesheet, err := h.trackUsecase.UpdateNotesheetFilepath(uint(notesheetID), uint(userID), filename)
	if err != nil {

		os.Remove(filepath)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notesheet)
}

// DownloadNotesheetFile handles GET /api/track/notesheet/file/{notesheetId}/{userId}
// Serves notesheet file download.
func (h *TrackHandler) DownloadNotesheetFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	notesheetID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid notesheet ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	notesheet, err := h.trackUsecase.GetNotesheet(uint(notesheetID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if notesheet.Filepath == "" {
		http.Error(w, "No file uploaded for this notesheet", http.StatusNotFound)
		return
	}

	filepath := path.Join(UPLOAD_DIR, notesheet.Filepath)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filepath)
}

// CreateNotesheetWithFile handles POST /api/track/notesheet/create
// Creates a new notesheet with file upload in one operation.
func (h *TrackHandler) CreateNotesheetWithFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil { // Poprawiony błąd składni
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	trackID, err := strconv.ParseUint(r.FormValue("track_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(r.FormValue("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var subgroupIDs []uint
	if subgroupIDsStr := r.FormValue("subgroup_ids"); subgroupIDsStr != "" {
		if err := json.Unmarshal([]byte(subgroupIDsStr), &subgroupIDs); err != nil {
			http.Error(w, "Invalid subgroup IDs format", http.StatusBadRequest)
			return
		}
	}

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), handler.Filename)
	filepath := path.Join(UPLOAD_DIR, filename)

	if err := os.MkdirAll(UPLOAD_DIR, 0755); err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(filepath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(filepath)
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	notesheet, err := h.trackUsecase.AddNotesheet(
		uint(trackID),
		handler.Header.Get("Content-Type"),
		filename,
		subgroupIDs,
		uint(userID),
	)

	if err != nil {
		os.Remove(filepath)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(notesheet)
}

// DeleteTrack handles DELETE /api/track/delete/{trackId}/{userId}
// Removes a track and its associated resources.
func (h *TrackHandler) DeleteTrack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	trackID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.trackUsecase.DeleteTrack(uint(trackID), uint(userID))
	if err != nil {
		if err.Error() == "access denied" || err.Error() == "insufficient permissions" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
