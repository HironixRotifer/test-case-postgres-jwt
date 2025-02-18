package httpserver

import (
	"net/http"

	authHandler "github.com/HironixRotifer/test-case-postgres-jwt/internal/http/handlers/auth"
	handlerPing "github.com/HironixRotifer/test-case-postgres-jwt/internal/http/handlers/ping"
	handlerSwagger "github.com/HironixRotifer/test-case-postgres-jwt/internal/http/handlers/swagger"

	"github.com/gorilla/mux"
)

func (s *ServerHTTP) initRoutes() *mux.Router {
	router := mux.NewRouter()

	// tokens
	router.HandleFunc("/api/v1/refresh", authHandler.RefreshTokens(s.provider.JWTCustomService())).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/access", authHandler.Auth(s.provider.JWTCustomService())).Methods(http.MethodGet)

	// swagger
	router.PathPrefix("/swagger/").Handler(handlerSwagger.Swagger()).Methods(http.MethodGet)
	router.HandleFunc("/ping", handlerPing.Ping()).Methods(http.MethodGet)

	return router
}
