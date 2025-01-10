package repositories_test

import (
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// SubgroupRepository Tests
func TestCreateSubgroup(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		subgroup      *model.Subgroup
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful subgroup creation",
			subgroup: &model.Subgroup{
				Name:        "Test Subgroup",
				GroupID:     1,
				Description: "Test Description",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "subgroups"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewSubgroupRepository()
			err := repo.CreateSubgroup(tc.subgroup)
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
