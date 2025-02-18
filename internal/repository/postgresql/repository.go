package repository

import "github.com/HironixRotifer/test-case-postgres-jwt/internal/models"

type JWTCustomRepository interface {
	GetUserByID(id int) (models.User, error)
	UpdateUserByID(id int, user models.User) error
}
