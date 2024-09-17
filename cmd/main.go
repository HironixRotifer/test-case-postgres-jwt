package main

import (
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	httpserve "github.com/HironixRotifer/test-case-postgres-jwt/internal/http"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/storage/postgres"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.MustLoadPath(".env")
	db, err := postgres.New(cfg)
	if err != nil {
		log.Err(err).Msg("error start database")
	}

	_, err = db.GetUserByID(1)
	log.Info().Msgf("%v", err)

	server := httpserve.NewServe(db)
	server.Start(cfg.Host, httpserve.WithPort(cfg.Port))
}
