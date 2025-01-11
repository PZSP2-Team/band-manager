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

func (h *TrackHandler) UploadNotesheetFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Pobierz ID notesheetu z URL
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

	// Parsuj multipart form
	r.ParseMultipartForm(10 << 20) // 10MB limit

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Utwórz unikalną nazwę pliku
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), handler.Filename)
	filepath := path.Join(UPLOAD_DIR, filename)

	// Upewnij się, że katalog istnieje
	if err := os.MkdirAll(UPLOAD_DIR, 0755); err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Utwórz nowy plik
	dst, err := os.Create(filepath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Skopiuj zawartość
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Zaktualizuj ścieżkę w bazie
	notesheet, err := h.trackUsecase.UpdateNotesheetFilepath(uint(notesheetID), uint(userID), filename)
	if err != nil {
		// Usuń plik jeśli nie udało się zaktualizować bazy
		os.Remove(filepath)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notesheet)
}

func (h *TrackHandler) DownloadNotesheetFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Pobierz ID notesheetu i użytkownika z URL
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

	// Pobierz notesheet i sprawdź uprawnienia
	notesheet, err := h.trackUsecase.GetNotesheet(uint(notesheetID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sprawdź czy ścieżka do pliku istnieje
	if notesheet.Filepath == "" {
		http.Error(w, "No file uploaded for this notesheet", http.StatusNotFound)
		return
	}

	filepath := path.Join(UPLOAD_DIR, notesheet.Filepath)

	// Sprawdź czy plik fizycznie istnieje
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Wyślij plik
	http.ServeFile(w, r, filepath)
}
