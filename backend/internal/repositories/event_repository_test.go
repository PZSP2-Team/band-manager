package repositories_test

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// EventRepository Tests
func TestCreateEvent(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	eventDate := time.Now()
	tests := []struct {
		name          string
		event         *model.Event
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful event creation",
			event: &model.Event{
				Title:       "Test Event",
				Location:    "Test Location",
				Description: "Test Description",
				Date:        eventDate,
				GroupID:     1,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "events"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewEventRepository()
			err := repo.CreateEvent(tc.event)
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
