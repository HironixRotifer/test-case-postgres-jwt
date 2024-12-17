package http

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/rs/zerolog/log"
	// "github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/middleware"
)

const (
	shutDownTimeout = 10 * time.Second
)

type ServerHTTP struct {
	port   int
	server *http.Server

	Router *gin.Engine
}

func NewServerHTTP(port int) *ServerHTTP {
	r := gin.Default()
	addr := fmt.Sprintf(":%v", port)

	server := &http.Server{Addr: addr, Handler: r}

	return &ServerHTTP{
		port:   port,
		server: server,
		Router: r,
	}
}

func (h *ServerHTTP) Start() {
	go func() {
		if err := h.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal().Err(err).Msg("failed to start http server")
			}
		}
	}()

	log.Info().Msg("http server is started!")
}

func (h *ServerHTTP) Stop(wg *sync.WaitGroup) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), shutDownTimeout)
		defer cancel()
		defer wg.Done()

		if err := h.server.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("Server forced to shutdown")
		} else {
			log.Info().Msg("HTTP stopped")
		}
	}()
}
