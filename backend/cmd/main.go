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

// enableCORS adds CORS headers to all HTTP responses.
// It configures allowed origins based on environment variables.
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

// main initializes the application, sets up services,
// configures HTTP routes, and starts the server.
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
	emailService := services.NewEmailService()
	if err != nil {
		log.Printf("Warning: Failed to initialize Google Calendar service: %v", err)

	}

	db.InitDB()

	authHandler := handlers.NewAuthHandler()
	groupHandler := handlers.NewGroupHandler()
	subgroupHandler := handlers.NewSubgroupHandler()
	trackHandler := handlers.NewTrackHandler()
	eventHandler := handlers.NewEventHandler(gcService, emailService)
	announcementHandler := handlers.NewAnnouncementHandler()
	adminHandler := handlers.NewAdminHandler()

	http.HandleFunc("/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}))
	// Authentication endpoints
	// POST /api/verify/login - Authenticates user and returns session data
	http.HandleFunc("/api/verify/login", enableCORS(authHandler.Login))
	// POST /api/verify/register - Creates new user account
	http.HandleFunc("/api/verify/register", enableCORS(authHandler.Register))

	// Group management endpoints
	// POST /api/group/create - Creates new band group
	// POST /api/group/join - Joins existing group using access token
	// GET /api/group/{groupId}/{userId} - Gets group details
	// GET /api/group/user/{userId} - Gets user's groups
	// GET /api/group/members/{groupId}/{userId} - Gets group members
	// DELETE /api/group/remove/{groupId}/{userId}/{requesterId} - Removes member from group
	// PUT /api/group/role/{groupId}/{userId}/{requesterId} - Updates member's role
	http.HandleFunc("/api/group/create", enableCORS(groupHandler.Create))
	http.HandleFunc("/api/group/join", enableCORS(groupHandler.Join))
	http.HandleFunc("/api/group/", enableCORS(groupHandler.GetGroupInfo))
	http.HandleFunc("/api/group/user/", enableCORS(groupHandler.GetUserGroups))
	http.HandleFunc("/api/group/members/", enableCORS(groupHandler.GetGroupMembers))
	http.HandleFunc("/api/group/remove/", enableCORS(groupHandler.RemoveMember))
	http.HandleFunc("/api/group/role/", enableCORS(groupHandler.UpdateMemberRole))

	// Subgroup management endpoints
	// POST /api/subgroup/create - Creates new subgroup
	// GET /api/subgroup/info/{subgroupId}/{userId} - Gets subgroup details
	// PUT /api/subgroup/update/{subgroupId}/{userId} - Updates subgroup
	// DELETE /api/subgroup/delete/{subgroupId}/{userId} - Deletes subgroup
	// POST /api/subgroup/members/add/{subgroupId}/{userId} - Adds members to subgroup
	// DELETE /api/subgroup/members/remove/{subgroupId}/{memberId}/{requesterId} - Removes member
	// GET /api/subgroup/group/{groupId}/{userId} - Gets all subgroups in group
	http.HandleFunc("/api/subgroup/create", enableCORS(subgroupHandler.Create))
	http.HandleFunc("/api/subgroup/info/", enableCORS(subgroupHandler.GetInfo))
	http.HandleFunc("/api/subgroup/update/", enableCORS(subgroupHandler.Update))
	http.HandleFunc("/api/subgroup/delete/", enableCORS(subgroupHandler.Delete))
	http.HandleFunc("/api/subgroup/members/add/", enableCORS(subgroupHandler.AddMembers))
	http.HandleFunc("/api/subgroup/members/remove/", enableCORS(subgroupHandler.RemoveMember))
	http.HandleFunc("/api/subgroup/group/", enableCORS(subgroupHandler.GetGroupSubgroups))

	// Track and notesheet management endpoints
	// POST /api/track/create - Creates new track
	// POST /api/track/notesheet - Adds notesheet to track
	// GET /api/track/user/notesheets/{trackId}/{userId} - Gets user's notesheets
	// GET /api/track/group/{groupId}/{userId} - Gets group's tracks
	// GET /api/track/notesheets/{trackId}/{userId} - Gets track's notesheets
	// POST /api/track/notesheet/upload/{notesheetId}/{userId} - Uploads notesheet file
	// GET /api/track/notesheet/file/{notesheetId}/{userId} - Downloads notesheet file
	// POST /api/track/notesheet/create - Creates notesheet with file
	// DELETE /api/track/delete/{trackId}/{userId} - Deletes track
	http.HandleFunc("/api/track/create", enableCORS(trackHandler.Create))
	http.HandleFunc("/api/track/notesheet", enableCORS(trackHandler.AddNotesheet))
	http.HandleFunc("/api/track/user/notesheets/", enableCORS(trackHandler.GetUserNotesheets))
	http.HandleFunc("/api/track/group/", enableCORS(trackHandler.GetGroupTracks))
	http.HandleFunc("/api/track/notesheets/", enableCORS(trackHandler.GetTrackNotesheets))
	http.HandleFunc("/api/track/notesheet/upload/", enableCORS(trackHandler.UploadNotesheetFile))
	http.HandleFunc("/api/track/notesheet/file/", enableCORS(trackHandler.DownloadNotesheetFile))
	http.HandleFunc("/api/track/notesheet/create/", enableCORS(trackHandler.CreateNotesheetWithFile))
	http.HandleFunc("/api/track/delete/", trackHandler.DeleteTrack)

	// Event management endpoints
	// POST /api/event/create - Creates new event
	// GET /api/event/info/{eventId}/{userId} - Gets event details
	// PUT /api/event/update/{eventId}/{userId} - Updates event
	// DELETE /api/event/delete/{eventId}/{userId} - Deletes event
	// GET /api/event/group/{groupId}/{userId} - Gets group's events
	// GET /api/event/user/{userId} - Gets user's events
	http.HandleFunc("/api/event/create", enableCORS(eventHandler.Create))
	http.HandleFunc("/api/event/info/", enableCORS(eventHandler.GetInfo))
	http.HandleFunc("/api/event/update/", enableCORS(eventHandler.Update))
	http.HandleFunc("/api/event/delete/", enableCORS(eventHandler.Delete))
	http.HandleFunc("/api/event/group/", enableCORS(eventHandler.GetGroupEvents))
	http.HandleFunc("/api/event/user/", enableCORS(eventHandler.GetUserEvents))

	// Announcement management endpoints
	// POST /api/announcement/create - Creates new announcement
	// DELETE /api/announcement/delete/{announcementId}/{userId} - Deletes announcement
	// GET /api/announcement/user/{userId} - Gets user's announcements
	// GET /api/announcement/group/{groupId}/{userId} - Gets group's announcements
	http.HandleFunc("/api/announcement/create", enableCORS(announcementHandler.Create))
	http.HandleFunc("/api/announcement/delete/", enableCORS(announcementHandler.Delete))
	http.HandleFunc("/api/announcement/user/", enableCORS(announcementHandler.GetUserAnnouncements))
	http.HandleFunc("/api/announcement/group/", enableCORS(announcementHandler.GetGroupAnnouncements))

	// Admin endpoints
	// PUT /api/admin/users/reset-password/{userId} - Resets user password
	// GET /api/admin/stats - Gets system statistics
	http.HandleFunc("/api/admin/users/reset-password/", enableCORS(adminHandler.ResetUserPassword))
	http.HandleFunc("/api/admin/stats", enableCORS(adminHandler.GetSystemStats))

	// Google Calendar integration endpoints
	// GET /api/calendar/auth - Initiates OAuth flow
	// GET /api/calendar/callback - Handles OAuth callback
	http.HandleFunc("/api/calendar/auth", enableCORS(eventHandler.GoogleCalendarAuth))
	http.HandleFunc("/api/calendar/callback", enableCORS(eventHandler.GoogleCalendarCallback))
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
