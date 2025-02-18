package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/jwt"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
	// redisService "github.com/HironixRotifer/test-case-postgres-jwt/internal/service/redisclient"
)

var (
	errUnauthorized = fmt.Errorf("Unauthorized")
)

func (o *service) RefreshToken(ctx context.Context, jwtCustom *models.JwtCustom) (*models.JwtCustom, error) {
	duration, err := time.ParseDuration(o.settings.REDIS_DURATION_SECOND)
	if err != nil {
		return nil, err
	}

	claims, err := jwt.ExtractTokenMetadata(jwtCustom.AccessToken)
	if err != nil {
		return nil, err
	}

	if err := claims.Valid(); err != nil {
		return nil, err
	}

	uidStr := strconv.Itoa(claims.Uid)

	result, err := o.clients.Get(ctx, uidStr)
	if err != nil {
		return nil, err
	}

	if result.RefreshToken != jwtCustom.RefreshToken {
		return nil, errUnauthorized
	}

	accessToken, refreshToken, err := jwt.GenerateJWTokens(claims.Uid)
	if err != nil {
		return nil, err
	}

	jwtCustom.AccessToken = accessToken
	jwtCustom.RefreshToken = refreshToken

	err = o.clients.Set(ctx, uidStr, jwtCustom, duration)
	if err != nil {
		return nil, err
	}

	return jwtCustom, nil
}
