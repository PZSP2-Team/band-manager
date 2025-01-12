package usecases

import (
	"band-manager-backend/internal/repositories"
	"band-manager-backend/internal/services"
	"context"
	"fmt"
)

type GoogleAuthUseCase struct {
	calendarService *services.GoogleCalendarService
	tokenRepo       *repositories.TokenRepository
}

func NewGoogleAuthUseCase() (*GoogleAuthUseCase, error) {
	calendarService, err := services.NewGoogleCalendarService()
	if err != nil {
		return nil, fmt.Errorf("failed to create calendar service: %v", err)
	}

	return &GoogleAuthUseCase{
		calendarService: calendarService,
		tokenRepo:       repositories.NewTokenRepository(),
	}, nil
}

func (u *GoogleAuthUseCase) HandleGoogleAuth(ctx context.Context, userID uint, authCode string) error {
	token, err := u.calendarService.ExchangeCodeForToken(ctx, authCode)
	if err != nil {
		return err
	}

	return u.tokenRepo.SaveTokens(userID, token.AccessToken, token.RefreshToken, token.Expiry)
}
