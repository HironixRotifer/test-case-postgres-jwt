package auth

import (
	def "github.com/HironixRotifer/test-case-postgres-jwt/internal/repository/redisdb"
	redis "github.com/redis/go-redis/v9"
)

var _ def.AuthRedisRepository = (*repository)(nil)

type repository struct {
	dbRedis *redis.Client
}

func NewRepository(dbRedis *redis.Client) *repository {
	return &repository{dbRedis: dbRedis}
}
