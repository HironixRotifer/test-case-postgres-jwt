package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	log "github.com/rs/zerolog/log"
)

// New создаёт новый клиент базы данных
func New(config *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%v user=%s password=%s  dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("error open database driver")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msg("error connection to database")
		return nil, err
	}

	log.Info().Msg("database connection is success")

	return db, nil
}
