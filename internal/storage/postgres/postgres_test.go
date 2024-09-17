package postgres

import (
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserByID(t *testing.T) {
	uid := 1
	expectedUser := models.User{
		UID:          uid,
		Email:        "test@example.com",
		IP:           "192.168.1.1",
		RefreshToken: "NewRefreshToken",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &Storage{db}

	mock.ExpectPrepare("UPDATE users SET refresh_token = \\$1 WHERE uid = \\$2")
	mock.ExpectExec("UPDATE users SET refresh_token = \\$1 WHERE uid = \\$2").
		WithArgs("NewRefreshToken", uid).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = s.updateUserByID(context.Background(), uid, expectedUser)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByID(t *testing.T) {
	uid := 1
	expectedUser := models.User{
		UID:          uid,
		Email:        "test@example.com",
		IP:           "192.168.1.1",
		RefreshToken: "refreshToken",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &Storage{db}

	mock.ExpectPrepare("SELECT \\* FROM users WHERE uid = \\$1")
	mock.ExpectQuery("SELECT \\* FROM users WHERE uid = \\$1").
		WithArgs(uid).
		WillReturnRows(sqlmock.NewRows([]string{"uid", "email", "ip", "refresh_token"}).
			AddRow(expectedUser.UID, expectedUser.Email, expectedUser.IP, expectedUser.RefreshToken))

	user, err := s.getUserByID(context.Background(), uid)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
