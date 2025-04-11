package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pull *pgxpool.Pool
}

func New() (*Database, error) {
	connString, err := genConnectionString()
	if err != nil {
		return nil, err
	}

	pull, err := pgxpool.New(context.Background(), connString)
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

// GetEncoded implements services.Storage.
func (db *Database) GetEncoded(short string) (string, error) {
	panic("unimplemented")
}

// GetOrdinary implements services.Storage.
func (db *Database) GetOrdinary(url string) (string, error) {
	panic("unimplemented")
}

// InsertURL implements services.Storage.
func (db *Database) InsertURL(URL string, short string) error {
	panic("unimplemented")
}
