package httpserver

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/provider"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/postgresql"
	log "github.com/rs/zerolog/log"
)

const (
	shutDownTimeout = 10 * time.Second
)

type ServerHTTP struct {
	server   http.Server
	provider *provider.Provider
	config   *config.Config
}

func NewServer(ctx context.Context) *ServerHTTP {
	server := &ServerHTTP{}

	err := server.initDependencies(ctx)
	if err != nil {
		panic(err)
	}

	return server
}

func (h *ServerHTTP) Run() {
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

func (h *ServerHTTP) initDependencies(ctx context.Context) error {
	deps := []func(context.Context) error{
		h.initConfig,
		h.initProvider,
		h.initServerHTTP,
	}

	for _, d := range deps {
		if err := d(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (h *ServerHTTP) initConfig(_ context.Context) error {
	h.config = config.MustLoadPath(".env")

	return nil
}

func (h *ServerHTTP) initProvider(_ context.Context) error {
	postgresDriver, err := postgresql.New(h.config)
	if err != nil {
		return err
	}

	provider.NewProvider(postgresDriver)

	return nil
}

func (h *ServerHTTP) initServerHTTP(ctx context.Context) error {
	h.server = http.Server{
		Addr: h.config.Host,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
		Handler: h.initRoutes(),
	}

	return nil
}
