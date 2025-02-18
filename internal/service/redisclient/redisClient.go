package clients

import (
	"context"
	"encoding/json"
	"time"

	settings "github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
	"github.com/redis/go-redis/v9"
)

type RedisClientDB interface {
	Set(ctx context.Context, userId string, jwt *models.JwtCustom, duration time.Duration) error
	Get(ctx context.Context, userId string) (jwt *models.JwtCustom, err error)
}

type RedisClientDBImp struct {
	client   *redis.Client
	settings *settings.Config
}

func NewRedisClient(client *redis.Client, settings *settings.Config) RedisClientDB {
	return &RedisClientDBImp{
		client:   client,
		settings: settings,
	}
}

func (r *RedisClientDBImp) Set(ctx context.Context, userId string, jwt *models.JwtCustom, duration time.Duration) error {
	jwtJSON, err := json.Marshal(jwt)
	if err != nil {
		return err
	}

	result := r.client.Set(ctx, userId, jwtJSON, duration)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (r *RedisClientDBImp) Get(ctx context.Context, userId string) (*models.JwtCustom, error) {
	jwt := &models.JwtCustom{}

	result, err := r.client.Get(ctx, userId).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(result), jwt)
	if err != nil {
		return nil, err
	}

	return jwt, nil
}
