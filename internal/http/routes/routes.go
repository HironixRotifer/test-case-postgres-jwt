package routes

import (
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/handlers"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/storage/postgres"
	"github.com/gin-gonic/gin"
)

func Routes(incomingRoutes *gin.Engine, db *postgres.Storage) {
	userHandler := handlers.New(db)

	incomingRoutes.POST("api-v1/refresh", userHandler.RefreshTokensByID())
	incomingRoutes.POST("api-v1/tokens", userHandler.GetTokensByID())
}
