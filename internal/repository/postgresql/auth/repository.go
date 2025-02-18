package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
	def "github.com/HironixRotifer/test-case-postgres-jwt/internal/repository/postgresql"
	_ "github.com/lib/pq"
)

var _ def.JWTCustomRepository = (*repository)(nil)

type repository struct {
	dbDriver *sql.DB
}

func NewRepository(dbDriver *sql.DB) *repository {
	return &repository{dbDriver: dbDriver}
}

// GetUserByID возвращает пользователя по его UID
func (r *repository) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	user, err := r.getUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *repository) getUserByID(ctx context.Context, uid int) (models.User, error) {
	var u models.User

	stmt, err := r.dbDriver.Prepare("SELECT * FROM users WHERE uid = $1")
	if err != nil {
		return u, err
	}

	row := stmt.QueryRowContext(ctx, uid)
	err = row.Scan(&u.UID, &u.Email, &u.Login, &u.Password)
	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUserByID обновляет пользователя по его UID
func (r *repository) UpdateUserByID(id int, user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	err := r.updateUserByID(ctx, id, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) updateUserByID(ctx context.Context, uid int, user models.User) error {
	stmt, err := r.dbDriver.Prepare("UPDATE users SET refresh_token = $1 WHERE uid = $2")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}
