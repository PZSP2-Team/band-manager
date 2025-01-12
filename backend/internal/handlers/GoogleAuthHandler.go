package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"log"
	"net/http"
)

type GoogleAuthHandler struct {
	googleAuthUseCase *usecases.GoogleAuthUseCase
}

func NewGoogleAuthHandler(googleAuthUseCase *usecases.GoogleAuthUseCase) *GoogleAuthHandler {
	return &GoogleAuthHandler{
		googleAuthUseCase: googleAuthUseCase,
	}
}

func (h *GoogleAuthHandler) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AuthCode string `json:"auth_code"`
		UserID   uint   `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.googleAuthUseCase.HandleGoogleAuth(r.Context(), request.UserID, request.AuthCode)
	if err != nil {
		http.Error(w, "Failed to handle Google auth", http.StatusInternalServerError)
		log.Printf("Google auth error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully connected with Google Calendar",
	})
}
