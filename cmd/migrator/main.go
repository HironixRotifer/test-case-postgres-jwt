package main

import (
	"errors"
	"fmt"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.MustLoadPath(".env")

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.DBHost, cfg.DBPort, cfg.DBName),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info().Msg("no migrations to apply")

			return
		}

		panic(err)
	}

	defer m.Close()

	log.Info().Msg("migrations applied")
}
