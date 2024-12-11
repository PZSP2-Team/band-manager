package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"fmt"
	"net/http"
)

type GroupHandler struct {
	groupUsecase *usecases.GroupUsecase
}

func NewGroupHandler() *GroupHandler {
	groupUsecase := usecases.NewGroupUsecase()
	return &GroupHandler{
		groupUsecase: groupUsecase,
	}
}

func (h *GroupHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request to /api/group/create")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from context (you need to implement authentication middleware)
	// userID := r.Context().Value("userID").(uint)
	userID := uint(9)

	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userRole, groupID, err := h.groupUsecase.CreateGroup(request.Name, request.Description, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"user_role":     userRole,
		"user_group_id": groupID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *GroupHandler) Join(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		UserID      uint   `json:"userId"`
		AccessToken string `json:"accessToken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userRole, groupID, err := h.groupUsecase.JoinGroup(request.UserID, request.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"ack":           true,
		"user_role":     userRole,
		"user_group_id": groupID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
