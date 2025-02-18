package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	httpserve "github.com/HironixRotifer/test-case-postgres-jwt/internal/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var wg sync.WaitGroup

func main() {
	ctx := context.Background()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	server := httpserve.NewServer(ctx)
	server.Run()

	<-signalChan
	log.Info().Msg("server graceful stop started")

	wg.Add(1)
	server.Stop(&wg)
	wg.Wait()
}
