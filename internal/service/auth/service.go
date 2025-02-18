package auth

import (
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	repository "github.com/HironixRotifer/test-case-postgres-jwt/internal/repository/postgresql"
	def "github.com/HironixRotifer/test-case-postgres-jwt/internal/service"
	clients "github.com/HironixRotifer/test-case-postgres-jwt/internal/service/redisclient"
)

var _ def.JWTCustomService = (*service)(nil)

type service struct {
	jwtCustomRepository repository.JWTCustomRepository
	clients             clients.RedisClientDB
	settings            *config.Config
}

func NewService(
	jwtCustomRepository repository.JWTCustomRepository,
	clients clients.RedisClientDB,
	config *config.Config,
) *service {
	return &service{
		clients:             clients,
		jwtCustomRepository: jwtCustomRepository,
		settings:            config,
	}
}
