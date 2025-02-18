package service

import (
	"context"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
)

type JWTCustomService interface {
	GetAccessToken(ctx context.Context, login, password string, uid int) (*models.JwtCustom, error)
	RefreshToken(ctx context.Context, jwtCustom *models.JwtCustom) (*models.JwtCustom, error)
}
