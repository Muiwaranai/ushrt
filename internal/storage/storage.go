package storage

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
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
		pool: pull,
	}, nil
}

func (db *Database) Close() {
	db.pool.Close()
}

func genConnectionString() (string, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" { // There is no password required in my app cuz its local
		return "", errors.New("Empty .env variable")
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName), nil
}

func (db *Database) ByShort(short string) (string, error) {
	var url string
	row := db.pool.QueryRow(context.Background(), "SELECT long_url FROM urls WHERE short_url=$1", short)
	err := row.Scan(&url)
	return url, err
}

func (db *Database) ByURL(url string) (string, error) {
	var short string
	row := db.pool.QueryRow(context.Background(), "SELECT short_url FROM urls WHERE long_url=$1", url)
	err := row.Scan(&short)
	return short, err
}

func (db *Database) Insert(url string, short string) error {
	_, err := db.pool.Exec(context.Background(), "INSERT INTO urls (long_url, short_url) VALUES ($1, $2)", url, short)
	return err
}
