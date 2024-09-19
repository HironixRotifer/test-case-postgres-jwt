package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/config"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Storage struct {
	db *sql.DB
}

// New создаёт новый клиент базы данных
func New(config *config.Config) (*Storage, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%v user=%s password=%s  dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Info().Msg("database connection is success")

	return &Storage{db}, nil
}

// GetUserByID возвращает пользователя по его UID
func (s *Storage) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	user, err := s.getUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *Storage) getUserByID(ctx context.Context, uid int) (models.User, error) {
	var u models.User

	stmt, err := s.db.Prepare("SELECT * FROM users WHERE uid = $1")
	if err != nil {
		return u, err
	}

	row := stmt.QueryRowContext(ctx, uid)
	err = row.Scan(&u.UID, &u.Email, &u.IP, &u.RefreshToken)
	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUserByID обновляет пользователя по его UID
func (s *Storage) UpdateUserByID(id int, user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	err := s.updateUserByID(ctx, id, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) updateUserByID(ctx context.Context, uid int, user models.User) error {
	stmt, err := s.db.Prepare("UPDATE users SET refresh_token = $1 WHERE uid = $2")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user.RefreshToken, uid)
	if err != nil {
		return err
	}

	return nil
}
