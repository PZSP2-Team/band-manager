package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      uint   `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userRole, groupID, err := h.groupUsecase.CreateGroup(request.Name, request.Description, request.UserID)
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
		UserID      uint   `json:"user_id"`
		AccessToken string `json:"access_token"`
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
		"user_role":     userRole,
		"user_group_id": groupID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *GroupHandler) GetGroupInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 { // /api/group/{groupId}/{userId}
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

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

	name, description, accessToken, err := h.groupUsecase.GetGroupInfo(uint(userID), uint(groupID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"name":         name,
		"description":  description,
		"access_token": accessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *GroupHandler) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

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

	members, err := h.groupUsecase.GetGroupMembers(uint(groupID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"members": members,
	})
}
