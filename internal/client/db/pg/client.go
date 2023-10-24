package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/romanfomindev/microservices-auth/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string, timeout int) (db.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(timeout))
	defer func() {
		cancel()
	}()

	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
