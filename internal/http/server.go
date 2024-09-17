package http

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/rs/zerolog/log"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/routes"
	// "github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/middleware"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/storage/postgres"
)

const (
	defaultHTTPPort = 8080
)

type HTTPServer struct {
	Port   int
	router *gin.Engine
}

// NewServe создаёт новый сервер
func NewServe(db *postgres.Storage) *HTTPServer {
	router := gin.Default()

	// router.Use(middleware.Authenticate())

	routes.Routes(router, db)

	return &HTTPServer{
		router: router,
	}
}

// Start запускает http сервер
// Priority: opts > port
// Default value is 8080, or value == 0 > random port
func (srv *HTTPServer) Start(addr string, opts ...Option) error {
	var options options        // Создаетcя пустая структура options
	for _, opt := range opts { // Производится перебор всех входных опций
		err := opt(&options)
		if err != nil {
			log.Err(err).Msg("error starting server")
			return err
		}
	}

	var port int
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}

	log.Info().Msgf("http serve is started on %v", port)
	err := srv.router.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Err(err).Msg("error starting server")
		return err
	}

	return nil
}

type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func randomPort() int {
	return rand.Intn(9000)
}
