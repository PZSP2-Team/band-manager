package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	usecase *usecases.UserUsecase
}

func NewAuthHandler() *AuthHandler {
	authUsecase := usecases.NewAuthUsecase()
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}
// Login handles the /api/auth/login endpoint
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
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

	response, err := h.authUsecase.Login(request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
