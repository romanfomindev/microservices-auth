package url_protected

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	"github.com/romanfomindev/platform_common/pkg/db"
)

type UrlProtectedRepository struct {
	db db.Client
}

func NewUrlProtectedPepository(db db.Client) repositories.UrlsProtectedRepository {
	return &UrlProtectedRepository{db: db}
}

func (r *UrlProtectedRepository) GetByUrl(ctx context.Context, url string) (*models.UrlProtected, error) {
	var urlProtected models.UrlProtected
	sqlStatement := "SELECT id, url, roles FROM urls_protected WHERE url = $1"
	q := db.Query{
		Name:     "url_protected_repository.GetByUrl",
		QueryRaw: sqlStatement,
	}

	err := r.db.DB().ScanOneContext(ctx, &urlProtected, q, url)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorNoRows
		}

		return nil, err
	}
	return &urlProtected, nil
}
