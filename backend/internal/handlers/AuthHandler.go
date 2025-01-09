package handlers

import (
	"band-manager-backend/internal/usecases"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	authUsecase *usecases.AuthUsecase
}

func NewAuthHandler() *AuthHandler {
	authUsecase := usecases.NewAuthUsecase()
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Struktura odpowiedzi zawierająca również token
	type LoginResponse struct {
		ID        uint   `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Token     string `json:"token"` // Dodajemy pole na token
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

	// Zauważ, że teraz Login zwraca również token
	user, groups, token, err := h.authUsecase.Login(request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := LoginResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Token:     token, // Dodajemy token do odpowiedzi
		Groups: make([]struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Role string `json:"role"`
		}, len(groups)),
	}

	// Wypełniamy informacje o grupach
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
