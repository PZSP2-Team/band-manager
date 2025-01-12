package repositories

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/services/db"
	"gorm.io/gorm"
	"time"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		db: db.GetDB(),
	}
}

func (r *TokenRepository) SaveTokens(userID uint, accessToken, refreshToken string, expiryTime time.Time) error {
	var token model.GoogleCalendarToken
	result := r.db.Where("user_id = ?", userID).First(&token)

	if result.Error == gorm.ErrRecordNotFound {
		token = model.GoogleCalendarToken{
			UserID:       userID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiryTime:   expiryTime,
		}
		return r.db.Create(&token).Error
	}

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	token.ExpiryTime = expiryTime
	return r.db.Save(&token).Error
}

func (r *TokenRepository) GetTokensByUserID(userID uint) (*model.GoogleCalendarToken, error) {
	var token model.GoogleCalendarToken
	err := r.db.Where("user_id = ?", userID).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *TokenRepository) DeleteTokens(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.GoogleCalendarToken{}).Error
}

func (r *TokenRepository) UpdateAccessToken(userID uint, accessToken string, expiryTime time.Time) error {
	return r.db.Model(&model.GoogleCalendarToken{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"access_token": accessToken,
			"expiry_time":  expiryTime,
		}).Error
}
