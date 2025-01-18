// Package handlers provides HTTP request handlers for the band management application.
// It implements the presentation layer, handling incoming HTTP requests,
// request validation, and response formatting.
package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// AdminHandler handles administrative operations.
type AdminHandler struct {
	adminUsecase *usecases.AdminUsecase
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{
		adminUsecase: usecases.NewAdminUsecase(),
	}
}

// ResetUserPassword handles POST /api/admin/users/reset-password/{userId}
// Resets the password for a specified user.
func (h *AdminHandler) ResetUserPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.ParseUint(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var request struct {
		NewPassword string `json:"new_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.adminUsecase.ResetUserPassword(uint(userID), request.NewPassword); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetSystemStats handles GET /api/admin/stats
// Retrieves system-wide statistics including total users and groups.
func (h *AdminHandler) GetSystemStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats, err := h.adminUsecase.GetSystemStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
