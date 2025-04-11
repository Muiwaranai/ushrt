package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pull *pgxpool.Pool
}

func New() (*Database, error) {
	pull, err := pgxpool.New(context.Background(), genConnectionString())
	if err != nil {
		return nil, err
	}

	return &Database{
		pull: pull,
	}, nil
}

func (db *Database) Close() {
	db.pull.Close()
}
