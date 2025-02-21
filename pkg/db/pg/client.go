package pg

import (
	"context"
	"fmt"

	"github.com/goterekhov/platform_common/pkg/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgClient struct {
	masterDBC db.DB
}

func NewPgClient(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{
			dbc: dbc,
		},
	}, nil
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}
