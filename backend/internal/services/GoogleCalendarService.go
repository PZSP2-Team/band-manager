package services

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"os"
	"time"
)

type GoogleCalendarService struct {
	config *oauth2.Config
}

func NewGoogleCalendarService() (*GoogleCalendarService, error) {
	b, err := os.ReadFile("internal/services/credentials.json")
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

func (s *GoogleCalendarService) ExchangeCodeForToken(ctx context.Context, authCode string) (*oauth2.Token, error) {
	token, err := s.config.Exchange(ctx, authCode)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}
	return token, nil
}

func (s *GoogleCalendarService) CreateEvent(ctx context.Context, token *oauth2.Token, eventData struct {
	Title       string
	Location    string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}) error {
	client := s.config.Client(ctx, token)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to create calendar client: %v", err)
	}

	event := &calendar.Event{
		Summary:     eventData.Title,
		Location:    eventData.Location,
		Description: eventData.Description,
		Start: &calendar.EventDateTime{
			DateTime: eventData.StartTime.Format(time.RFC3339),
			TimeZone: "Europe/Warsaw",
		},
		End: &calendar.EventDateTime{
			DateTime: eventData.EndTime.Format(time.RFC3339),
			TimeZone: "Europe/Warsaw",
		},
	}

	calendarID := "primary"
	_, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		return fmt.Errorf("unable to create event: %v", err)
	}

	return nil
}
