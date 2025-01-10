package repositories_test

import (
	"band-manager-backend/internal/db"
	"band-manager-backend/internal/model"
	"band-manager-backend/internal/repositories"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func mockGetDB(mockDB *gorm.DB) func() {
	db.SetDB(mockDB)
	return func() {
		db.SetDB(nil)
	}
}

func newTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))
	require.NoError(t, err)

	return db, mock
}

func TestGetUserByEmail(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		email         string
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedUser  *model.User
		expectedError error
	}{
		{
			name:  "successful user fetch",
			email: "test@example.com",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password_hash"}).
					AddRow(1, "John", "Doe", "test@example.com", "hashedpassword")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`)).
					WithArgs("test@example.com", 1).
					WillReturnRows(rows)
			},
			expectedUser: &model.User{
				ID:           1,
				FirstName:    "John",
				LastName:     "Doe",
				Email:        "test@example.com",
				PasswordHash: "hashedpassword",
			},
			expectedError: nil,
		},
		{
			name:  "user not found",
			email: "notfound@example.com",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`)).
					WithArgs("notfound@example.com", 1).
					WillReturnError(sql.ErrNoRows)
			},
			expectedUser:  nil,
			expectedError: errors.New("użytkownik nie znaleziony"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewUserRepository()

			user, err := repo.GetUserByEmail(tc.email)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedUser, user)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		user          *model.User
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful user creation",
			user: &model.User{
				FirstName:    "John",
				LastName:     "Doe",
				Email:        "john@example.com",
				PasswordHash: "hashedpassword",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
		{
			name: "duplicate email error",
			user: &model.User{
				FirstName:    "John",
				LastName:     "Doe",
				Email:        "existing@example.com",
				PasswordHash: "hashedpassword",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
					WillReturnError(&pq.Error{Code: "23505"})
				mock.ExpectRollback()
			},
			expectedError: errors.New("nie udało się utworzyć użytkownika"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewUserRepository()

			err := repo.CreateUser(tc.user)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		user          *model.User
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful update",
			user: &model.User{
				ID:           1,
				FirstName:    "John",
				LastName:     "Doe Updated",
				Email:        "john@example.com",
				PasswordHash: "hashedpassword",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewUserRepository()

			err := repo.UpdateUser(tc.user)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	mockDB, mock := newTestDB(t)
	cleanup := mockGetDB(mockDB)
	defer cleanup()

	tests := []struct {
		name          string
		userID        uint
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:   "successful deletion",
			userID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users"`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mock)
			repo := repositories.NewUserRepository()

			err := repo.DeleteUser(tc.userID)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
