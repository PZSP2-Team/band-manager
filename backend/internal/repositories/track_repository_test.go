package repositories_test

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// TrackRepository Tests
func TestCreateTrack(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		track         *model.Track
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful track creation",
			track: &model.Track{
				Name:        "Test Track",
				GroupID:     1,
				Description: "Test Description",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tracks"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewTrackRepository()
			err := repo.CreateTrack(tc.track)
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetGroupTracks(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		groupID       uint
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:    "successful tracks fetch",
			groupID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "group_id", "description"}).
					AddRow(1, "Track 1", 1, "Description 1").
					AddRow(2, "Track 2", 1, "Description 2")

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tracks"`)).
					WithArgs(1).
					WillReturnRows(rows)

				// Dla Preload Notesheets
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "notesheets"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}))
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewTrackRepository()
			tracks, err := repo.GetGroupTracks(tc.groupID)
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, tracks)
				assert.Len(t, tracks, 2)
			}
		})
	}
}
