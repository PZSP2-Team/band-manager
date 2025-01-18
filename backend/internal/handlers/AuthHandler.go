// Package handlers provides HTTP request handlers for the band management application.
// It implements the presentation layer, handling incoming HTTP requests,
// request validation, and response formatting.
package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"net/http"
)

// AuthHandler manages user authentication and registration.
type AuthHandler struct {
	authUsecase *usecases.AuthUsecase
}

func NewAuthHandler() *AuthHandler {
	authUsecase := usecases.NewAuthUsecase()
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

// Login handles POST /api/verify/login
// Authenticates user and returns user details with group memberships.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	type LoginResponse struct {
		ID        uint   `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Groups    []struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Role string `json:"role"`
		} `json:"groups"`
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, groups, err := h.authUsecase.Login(request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := LoginResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Groups: make([]struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Role string `json:"role"`
		}, len(groups)),
	}

	for i, group := range groups {
		response.Groups[i] = struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Role string `json:"role"`
		}{
			ID:   group.ID,
			Name: group.Name,
			Role: group.Role,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Register handles POST /api/verify/register
// Creates a new user account with provided details.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.authUsecase.Register(request.FirstName, request.LastName, request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
}
