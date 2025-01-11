package services

import (
	"band-manager-backend/internal/config"
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"context"
	"fmt"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendarService struct {
	config *oauth2.Config
}

func (s *GoogleCalendarService) GetOrCreateToken() (*oauth2.Token, error) {
	var gToken model.GoogleToken
	result := db.GetDB().First(&gToken)
	if result.Error != nil {
		return nil, fmt.Errorf("no token found - authorization required")
	}
	return &oauth2.Token{
		AccessToken:  gToken.AccessToken,
		TokenType:    gToken.TokenType,
		RefreshToken: gToken.RefreshToken,
		Expiry:       gToken.Expiry,
	}, nil
}

func NewGoogleCalendarService(cfg *config.Config) (*GoogleCalendarService, error) {
	b, err := os.ReadFile(cfg.GoogleCalendarConfig.CredentialsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	return &GoogleCalendarService{
		config: config,
	}, nil
}

// Tu będą dodatkowe metody serwisu...

func (s *GoogleCalendarService) CreateCalendarEvent(event *model.Event) error {

	token, err := s.GetOrCreateToken()
	if err != nil {
		return err
	}

	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithTokenSource(
		s.config.TokenSource(ctx, token),
	))
	if err != nil {
		return fmt.Errorf("unable to create calendar client: %v", err)
	}

	calendarEvent := &calendar.Event{
		Summary:     event.Title,
		Location:    event.Location,
		Description: event.Description,
		Start: &calendar.EventDateTime{
			DateTime: event.Date.Format(time.RFC3339),
			TimeZone: "Europe/Warsaw",
		},
		End: &calendar.EventDateTime{
			DateTime: event.Date.Add(2 * time.Hour).Format(time.RFC3339), // Domyślnie 2h
			TimeZone: "Europe/Warsaw",
		},
	}

	createdEvent, err := srv.Events.Insert("primary", calendarEvent).Do()
	if err != nil {
		return fmt.Errorf("unable to create event in calendar: %v", err)
	}

	// Zapisz powiązanie w bazie
	gcEvent := &model.GoogleCalendarEvent{
		EventID:    event.ID,
		CalendarID: createdEvent.Id,
		LastSynced: time.Now(),
	}

	return db.GetDB().Create(gcEvent).Error
}

func (s *GoogleCalendarService) GetAuthURL() string {
	return s.config.AuthCodeURL("state")
}

func (s *GoogleCalendarService) SaveToken(code string) error {
	fmt.Println("Exchanging authorization code for token...") // DEBUG

	token, err := s.config.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging code for token:", err) // DEBUG
		return err
	}
	fmt.Printf("Token received: %+v\n", token) // DEBUG

	gToken := &model.GoogleToken{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
	fmt.Println("Saving token to database...") // DEBUG

	err = db.GetDB().Create(gToken).Error
	if err != nil {
		fmt.Println("Error saving token to database:", err) // DEBUG
	}
	return err
}
