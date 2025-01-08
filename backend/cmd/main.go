package main

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		frontendHost := os.Getenv("FRONTEND_HOST")
		if frontendHost == "" {
			frontendHost = "localhost"
		}

		frontendPort := os.Getenv("FRONTEND_PORT")
		if frontendPort == "" {
			frontendPort = "3000"
		}

		allowedOrigin := fmt.Sprintf("http://%s:%s", frontendHost, frontendPort)
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	db.InitDB()

	authHandler := handlers.NewAuthHandler()
	groupHandler := handlers.NewGroupHandler()

	http.HandleFunc("/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}))
	http.HandleFunc("/api/verify/login", enableCORS(authHandler.Login))
	http.HandleFunc("/api/verify/register", enableCORS(authHandler.Register))

	http.HandleFunc("/api/group/create", enableCORS(groupHandler.Create))
	http.HandleFunc("/api/group/join", enableCORS(groupHandler.Join))
	http.HandleFunc("/api/group/", enableCORS(groupHandler.GetGroupInfo))
	http.HandleFunc("/api/group/user/", enableCORS(groupHandler.GetUserGroups))
	http.HandleFunc("/api/group/members/", enableCORS(groupHandler.GetGroupMembers))
	http.HandleFunc("/api/group/remove/", enableCORS(groupHandler.RemoveMember))

	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
