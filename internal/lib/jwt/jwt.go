package jwt

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
	IP    string `json:"ip"`
	jwt.StandardClaims
}

func TokenGenerator(uid string) (accessToken string, refreshToken string, err error) {
	claims := &SignedDetails{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshclaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return
	}

	return token, refreshtoken, err
}

func ValidateToken(signedtoken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, fmt.Errorf("the token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, fmt.Errorf("token is already expired")
	}

	return claims, nil

}
