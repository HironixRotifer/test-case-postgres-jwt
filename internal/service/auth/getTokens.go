package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/jwt"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
)

var (
	errInvalidCredentials = fmt.Errorf("invalid login or password")
)

func (o *service) GetAccessToken(ctx context.Context, login, password string, uid int) (*models.JwtCustom, error) {
	jwtCustom := &models.JwtCustom{}
	uidStr := strconv.Itoa(uid)
	duration, err := time.ParseDuration(o.settings.REDIS_DURATION_SECOND)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user, err := o.jwtCustomRepository.GetUserByID(uid)
	if err != nil {
		return nil, err
	}

	if user.Login != login {
		return nil, errInvalidCredentials
	}

	accessToken, refreshToken, err := jwt.GenerateJWTokens(uid)
	if err != nil {
		return nil, err
	}

	jwtCustom.AccessToken = accessToken
	jwtCustom.RefreshToken = refreshToken

	err = o.clients.Set(ctx, uidStr, jwtCustom, duration)
	if err != nil {
		return nil, err
	}

	return jwtCustom, err
}
