package user

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user/convertor"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user/model"
	"github.com/romanfomindev/platform_common/pkg/db"
)

var _ repositories.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db db.Client
}

func NewUserRepository(db db.Client) repositories.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user model.UserCreate) (uint64, error) {
	var lastInsertId uint64

	sqlStatement := "INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id"
	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: sqlStatement,
	}

	err := r.db.DB().QueryRowContext(ctx, q, user.Name, user.Email, user.Password, user.Role).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (r *UserRepository) Update(ctx context.Context, id uint64, user model.UserUpdate) error {
	sqlStatement := "UPDATE users SET name = $1, email = $2,  role = $3 WHERE id = $4"
	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: sqlStatement,
	}

	_, err := r.db.DB().ExecContext(ctx, q, user.Name, user.Email, user.Role, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetById(ctx context.Context, id uint64) (*models.User, error) {
	var user model.User
	sqlStatement := "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = $1"
	q := db.Query{
		Name:     "user_repository.GetById",
		QueryRaw: sqlStatement,
	}

	err := r.db.DB().ScanOneContext(ctx, &user, q, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorNoRows
		}

		return nil, err
	}

	return convertor.ToUserFromUserRepo(user), nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	sqlStatement := "DELETE FROM users  WHERE id = $1"
	q := db.Query{
		Name:     "user_repository.DeleteById",
		QueryRaw: sqlStatement,
	}

	_, err := r.db.DB().ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user model.User
	sqlStatement := "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1"
	q := db.Query{
		Name:     "user_repository.GetById",
		QueryRaw: sqlStatement,
	}

	err := r.db.DB().ScanOneContext(ctx, &user, q, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorNoRows
		}

		return nil, err
	}

	return convertor.ToUserFromUserRepo(user), nil
}
