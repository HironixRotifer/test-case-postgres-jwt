package hashpassword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDoPasswordsMatch проверяет работу функции DoPasswordsMatch в ситуациях с верным и не верным паролями.
func TestDoPasswordsMatch(t *testing.T) {
	salt := GenerateRandomSalt()
	correctPassword := HashPassword("password123", salt)

	tests := []struct {
		name         string
		password     string
		hashpassword string
		expected     bool
	}{
		{
			name:         "password correct",
			password:     "password123",
			hashpassword: correctPassword,
			expected:     true,
		},
		{
			name:         "password incorrect",
			password:     "password123",
			hashpassword: "incorrecthash",
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := DoPasswordsMatch(tt.hashpassword, tt.password, salt)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// TestHashPassword проверяет корректность генерации хэшей.
func TestHashPassword(t *testing.T) {
	salt := GenerateRandomSalt()
	tests := []struct {
		name         string
		password     string
		hashpassword string
	}{
		{
			name:         "password correct",
			password:     "password123",
			hashpassword: HashPassword("password123", salt),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashPassword := HashPassword(tt.password, salt)
			assert.Equal(t, tt.hashpassword, hashPassword)
		})
	}
}
