package main

import (
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/handlers"
	// "github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.Use(middleware.Authenticate())
	router.Handle("POST", "api-v1/gettokens", handlers.GetTokensByID())
	router.Handle("GET", "api-v1/gettokens", handlers.RefreshTokensByID())

	// routes(router)

	router.Run()
}
