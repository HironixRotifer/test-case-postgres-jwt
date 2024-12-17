package routes

import (
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/handlers"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/storage/postgres"
	"github.com/gin-gonic/gin"
)

func InitRoutesWithStorage(routerGroup *gin.RouterGroup, db *postgres.Storage) {
	userHandler := handlers.New(db)

	routerGroup.POST("api-v1/refresh", userHandler.RefreshTokensByID())
	routerGroup.POST("api-v1/tokens", userHandler.GetTokensByID())
}
