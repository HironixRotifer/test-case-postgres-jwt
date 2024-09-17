package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenGenerator_SuccessCreateJWT(t *testing.T) {
	tests := []struct {
		name  string
		token string
		uid   int
		email string
		ip    string
	}{
		{
			name: "success generate",
			uid:  1,
			ip:   "192.168.0.1",
		},
		{
			name: "same data",
			uid:  2,
			ip:   "192.168.1.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := TokenGenerator(tt.uid, tt.ip)
			assert.NoError(t, err)

		})
	}
}

func TestTokenGenerator_NotSameValues(t *testing.T) {
	tests := []struct {
		name string
		uid  int
		ip   string
	}{
		{
			name: "same data",
			uid:  1,
			ip:   "192.168.0.1",
		},
		{
			name: "same data",
			uid:  1,
			ip:   "192.168.0.1",
		},
		{
			name: "same data",
			uid:  1,
			ip:   "192.168.0.1",
		},
	}

	var token string

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessToken, _, err := TokenGenerator(tt.uid, tt.ip)
			time.Sleep(1 * time.Millisecond)
			assert.NoError(t, err)
			assert.NotEqual(t, token, accessToken)
		})
	}
}

func TestValidateToken_DecodeAccessToken(t *testing.T) {
	tests := []struct {
		name string
		uid  int
		ip   string
	}{
		{
			name: "test claims",
			uid:  1,
			ip:   "192.168.0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessToken, _, err := TokenGenerator(tt.uid, tt.ip)
			assert.NoError(t, err)

			got, err := ValidateToken(accessToken)
			assert.NoError(t, err)

			assert.Equal(t, tt.uid, got.Uid)
			assert.Equal(t, tt.ip, got.IP)
		})
	}
}

func TestRequireTokens(t *testing.T) {
	tests := []struct {
		name  string
		uid   int
		email string
		ip    string
	}{
		{
			name:  "identical tokens",
			uid:   1,
			email: "example@mail.com",
			ip:    "192.168.0.1",
		},
		{
			name:  "different tokens",
			uid:   1,
			email: "example@mail.com",
			ip:    "192.168.1.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessToken, refreshToken, err := TokenGenerator(tt.uid, tt.ip)
			assert.NoError(t, err)

			err = RequireTokens(accessToken, refreshToken)
			assert.NoError(t, err)
		})
	}
}
