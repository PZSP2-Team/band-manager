package repositories_test

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// GroupRepository Tests
func TestCreateGroup(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		group         *model.Group
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful group creation",
			group: &model.Group{
				Name:        "Test Group",
				AccessToken: "test-token",
				Description: "Test Description",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "groups"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewGroupRepository()
			err := repo.CreateGroup(tc.group)
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetGroupByAccessToken(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		accessToken   string
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:        "group found",
			accessToken: "test-token",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "access_token", "description"}).
					AddRow(1, "Test Group", "test-token", "Description")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "groups"`)).
					WithArgs("test-token", 1).
					WillReturnRows(rows)
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewGroupRepository()
			group, err := repo.GetGroupByAccessToken(tc.accessToken)
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, group)
			}
		})
	}
}
