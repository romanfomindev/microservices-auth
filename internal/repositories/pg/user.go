package pg

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/romanfomindev/microservices-auth/internal/config"
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
)

var _ repositories.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(ctx context.Context, cfg config.PGConfig) (repositories.UserRepository, error) {
	conn, err := pgx.Connect(ctx, cfg.DSN())
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		conn: conn,
	}, nil
}

/** TODO наверно нада какая нить DTO */
func (r *UserRepository) Create(ctx context.Context, name, email, password, role string) (uint64, error) {
	var lastInsertId uint64

	sqlStatement := "INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id"

	err := r.conn.QueryRow(ctx, sqlStatement, name, email, password, role).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (r *UserRepository) Update(ctx context.Context, id uint64, name, email, role string) error {
	sqlStatement := "UPDATE users SET name = $1, email = $2,  role = $3 WHERE id = $4"
	_, err := r.conn.Exec(ctx, sqlStatement, name, email, role, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetById(ctx context.Context, id uint64) (models.User, error) {
	var user models.User
	sqlStatement := "SELECT name, email, password, role, created_at, updated_at FROM users WHERE id = $1"

	err := r.conn.QueryRow(ctx, sqlStatement, id).
		Scan(&user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	sqlStatement := "DELETE FROM users  WHERE id = $1"
	_, err := r.conn.Exec(ctx, sqlStatement, id)

	if err != nil {
		return err
	}

	return nil
}
