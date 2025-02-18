package provider

import (
	"database/sql"
	"fmt"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	authRepository "github.com/HironixRotifer/test-case-postgres-jwt/internal/repository/postgresql/auth"
	"github.com/redis/go-redis/v9"

	repository "github.com/HironixRotifer/test-case-postgres-jwt/internal/repository/postgresql"
	service "github.com/HironixRotifer/test-case-postgres-jwt/internal/service"
	authService "github.com/HironixRotifer/test-case-postgres-jwt/internal/service/auth"
	clients "github.com/HironixRotifer/test-case-postgres-jwt/internal/service/redisclient"
)

type Provider struct {
	dbDriver    *sql.DB
	settings    *config.Config
	clientRedis clients.RedisClientDB

	authRepository repository.JWTCustomRepository

	authService service.JWTCustomService
}

func NewProvider(db *sql.DB) *Provider {
	return &Provider{
		dbDriver: db,
	}
}

func (p *Provider) JWTCustomRepository() repository.JWTCustomRepository {
	if p.authRepository == nil {
		p.authRepository = authRepository.NewRepository(p.dbDriver)
	}

	return p.authRepository
}

func (p *Provider) JWTCustomService() service.JWTCustomService {
	if p.authService == nil {
		p.authService = authService.NewService(p.JWTCustomRepository(), p.NewRedisCLient(), p.NewSettings())
	}

	return p.authService
}

func (p *Provider) NewRedisCLient() clients.RedisClientDB {
	if p.clientRedis == nil {
		if p.settings == nil {
			p.settings = p.NewSettings()
		}
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", p.settings.REDIS_HOST, p.settings.REDIS_PORT),
			Password: p.settings.REDIS_PASSWORD,
			DB:       0,
		})

		p.clientRedis = clients.NewRedisClient(client, p.NewSettings())
	}

	return p.clientRedis
}

func (p *Provider) NewSettings() *config.Config {
	if p.settings == nil {
		p.settings = config.MustLoadPath(".env")
	}

	return p.settings
}
