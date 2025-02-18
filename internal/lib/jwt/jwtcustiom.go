package jwt

import (
	"fmt"
	"time"

	"github.com/HironixRotifer/test-case-postgres-jwt/pkg/generator"
	jwt "github.com/dgrijalva/jwt-go"
)

const secretKey = "sercret"

type AandaClaims struct {
	Uid int `json:"uid,omitempty"`
	Jti int `json:"jti,omitempty"`
	jwt.StandardClaims
}

// GenerateJWTokens генерирует пару access и refresh токенов
func GenerateJWTokens(uid int) (accessToken string, refreshToken string, err error) {

	claims := &AandaClaims{
		Uid: uid,
		Jti: generator.Jti(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}

	refreshclaims := &AandaClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

// ExtractTokenMetadata валuдирует и декодирует информацию из токена
func ExtractTokenMetadata(accessToken string) (claims *AandaClaims, err error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&AandaClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AandaClaims)
	if !ok {
		return nil, fmt.Errorf("the token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, fmt.Errorf("token is already expired")
	}

	return claims, nil
}
