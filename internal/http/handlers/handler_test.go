package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/jwt"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUserProvider struct {
	GetUserByIDFunc    func(id int) (models.User, error)
	UpdateUserByIDFunc func(id int, user models.User) error
}

func (m *MockUserProvider) GetUserByID(id int) (models.User, error) {
	return m.GetUserByIDFunc(id)
}

func (m *MockUserProvider) UpdateUserByID(id int, user models.User) error {
	return m.UpdateUserByIDFunc(id, user)
}

func TestGetTokensByID_successGetTokens(t *testing.T) {
	mockProvider := &MockUserProvider{
		GetUserByIDFunc: func(id int) (models.User, error) {
			return models.User{
				UID:          1,
				Email:        "example",
				IP:           "example@mail.com",
				RefreshToken: "",
			}, nil
		},
		UpdateUserByIDFunc: func(id int, user models.User) error {
			return nil
		},
	}

	requestBody, _ := json.Marshal(map[string]int{
		"guid": 1,
	})

	req, err := http.NewRequest("POST", "/tokens", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}

	u := New(mockProvider)
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/tokens", u.GetTokensByID())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}

func TestRefreshTokensByID(t *testing.T) {
	accessToken, refreshToken, err := jwt.TokenGenerator(1, "192.168.0.1")
	assert.NoError(t, err)

	mockProvider := &MockUserProvider{
		GetUserByIDFunc: func(id int) (models.User, error) {
			return models.User{
				UID:          1,
				Email:        "example@mail.com",
				IP:           "192.168.0.1",
				RefreshToken: refreshToken,
			}, nil
		},
		UpdateUserByIDFunc: func(id int, user models.User) error {
			return nil
		},
	}

	time.Sleep(1 * time.Millisecond)

	requestBody, _ := json.Marshal(map[string]string{
		"token":         accessToken,
		"refresh_token": refreshToken,
	})

	req, err := http.NewRequest("POST", "/refresh", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}

	u := New(mockProvider)
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/refresh", u.RefreshTokensByID())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}
