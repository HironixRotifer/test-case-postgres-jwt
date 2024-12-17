package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	httpserve "github.com/HironixRotifer/test-case-postgres-jwt/internal/http"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/routes"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/storage/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var wg sync.WaitGroup

func main() {
	cfg := config.MustLoadPath(".env")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	db, err := postgres.New(cfg)
	if err != nil {
		log.Error().Err(err).Msg("error start database")
	}

	server := httpserve.NewServerHTTP(cfg.Port)
	routWithStorage := server.Router.Group("/")
	routes.InitRoutesWithStorage(routWithStorage, db)

	server.Start()

	<-signalChan
	log.Info().Msg("server graceful stop started")

	wg.Add(1)
	server.Stop(&wg)
	wg.Wait()
}
