package main

import (
	"band-manager-backend/internal/config"
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/handlers" // Nowy import
	"band-manager-backend/internal/services"
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	gcService, err := services.NewGoogleCalendarService(cfg)
	if err != nil {
		log.Printf("Warning: Failed to initialize Google Calendar service: %v", err)
		// Nie przerywamy działania aplikacji, ale logujemy ostrzeżenie
	}

	db.InitDB()

	authHandler := handlers.NewAuthHandler()
	groupHandler := handlers.NewGroupHandler()
	subgroupHandler := handlers.NewSubgroupHandler()
	trackHandler := handlers.NewTrackHandler()
	eventHandler := handlers.NewEventHandler(gcService)
	announcementHandler := handlers.NewAnnouncementHandler()
	adminHandler := handlers.NewAdminHandler()

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
	http.HandleFunc("/api/group/role/", enableCORS(groupHandler.UpdateMemberRole))

	http.HandleFunc("/api/subgroup/create", enableCORS(subgroupHandler.Create))
	http.HandleFunc("/api/subgroup/info/", enableCORS(subgroupHandler.GetInfo))
	http.HandleFunc("/api/subgroup/update/", enableCORS(subgroupHandler.Update))
	http.HandleFunc("/api/subgroup/delete/", enableCORS(subgroupHandler.Delete))
	http.HandleFunc("/api/subgroup/members/add/", enableCORS(subgroupHandler.AddMembers))
	http.HandleFunc("/api/subgroup/members/remove/", enableCORS(subgroupHandler.RemoveMember))
	http.HandleFunc("/api/subgroup/group/", enableCORS(subgroupHandler.GetGroupSubgroups))

	http.HandleFunc("/api/track/create", enableCORS(trackHandler.Create))
	http.HandleFunc("/api/track/notesheet", enableCORS(trackHandler.AddNotesheet))
	http.HandleFunc("/api/track/user/notesheets/", enableCORS(trackHandler.GetUserNotesheets))
	http.HandleFunc("/api/track/group/", enableCORS(trackHandler.GetGroupTracks))
	http.HandleFunc("/api/track/notesheets/", enableCORS(trackHandler.GetTrackNotesheets))
	// Event endpoints
	http.HandleFunc("/api/event/create", enableCORS(eventHandler.Create))
	http.HandleFunc("/api/event/info/", enableCORS(eventHandler.GetInfo))
	http.HandleFunc("/api/event/update/", enableCORS(eventHandler.Update))
	http.HandleFunc("/api/event/delete/", enableCORS(eventHandler.Delete))
	http.HandleFunc("/api/event/group/", enableCORS(eventHandler.GetGroupEvents))
	http.HandleFunc("/api/event/user/", enableCORS(eventHandler.GetUserEvents))

	http.HandleFunc("/api/announcement/create", enableCORS(announcementHandler.Create))
	http.HandleFunc("/api/announcement/delete/", enableCORS(announcementHandler.Delete))
	http.HandleFunc("/api/announcement/user/", enableCORS(announcementHandler.GetUserAnnouncements))
	http.HandleFunc("/api/announcement/group/", enableCORS(announcementHandler.GetGroupAnnouncements))

	http.HandleFunc("/api/admin/users/reset-password/", enableCORS(adminHandler.ResetUserPassword))
	http.HandleFunc("/api/admin/stats", enableCORS(adminHandler.GetSystemStats))
	http.HandleFunc("/api/track/notesheet/upload/", enableCORS(trackHandler.UploadNotesheetFile))
	http.HandleFunc("/api/track/notesheet/file/", enableCORS(trackHandler.DownloadNotesheetFile))
	http.HandleFunc("/api/track/notesheet/create/", enableCORS(trackHandler.CreateNotesheetWithFile))
	http.HandleFunc("/api/track/delete/", trackHandler.DeleteTrack)

	http.HandleFunc("/api/calendar/auth", enableCORS(eventHandler.GoogleCalendarAuth))
	http.HandleFunc("/api/calendar/callback", enableCORS(eventHandler.GoogleCalendarCallback))
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
