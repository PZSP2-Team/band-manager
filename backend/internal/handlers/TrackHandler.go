package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type TrackHandler struct {
	trackUsecase *usecases.TrackUsecase
}

func NewTrackHandler() *TrackHandler {
	return &TrackHandler{
		trackUsecase: usecases.NewTrackUsecase(),
	}
}

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

func (h *TrackHandler) AddNotesheet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		TrackID     uint   `json:"track_id"`
		UserID      uint   `json:"user_id"`
		Filepath    string `json:"filepath"` // Ścieżka do pliku z frontendu
		Instrument  string `json:"instrument"`
		SubgroupIDs []uint `json:"subgroup_ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	notesheet, err := h.trackUsecase.AddNotesheet(
		request.TrackID,
		request.Instrument,
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

func (h *TrackHandler) GetSubgroupNotesheets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	subgroupID, err := strconv.ParseUint(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid subgroup ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	notesheets, err := h.trackUsecase.GetSubgroupNotesheets(uint(subgroupID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"notesheets": notesheets,
	})
}
