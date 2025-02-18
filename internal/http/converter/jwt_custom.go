package converter

import (
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/request"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/response"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
)

func ToResponseJwtCustomFromModel(ms *models.JwtCustom) *response.JwtCustom {
	jwtCusom := &response.JwtCustom{
		AccessToken:  ms.AccessToken,
		RefreshToken: ms.RefreshToken,
	}

	return jwtCusom
}

func ToModelJwtCustomFromRequest(ms *request.JwtCustom) *models.JwtCustom {
	jwtCusom := &models.JwtCustom{
		AccessToken:  ms.AccessToken,
		RefreshToken: ms.RefreshToken,
	}

	return jwtCusom
}
