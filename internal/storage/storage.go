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

func (db *Database) GetEncoded(url string) (string, error) {
	var encodedurl string
	err := db.pool.QueryRow(context.Background(), "SELECT encoded_url COUNT FROM urls WHERE original_url = $1", url).Scan(&encodedurl)
	if err != nil {
		return "", err
	}
	return encodedurl, nil
}

func (db *Database) GetOrdinary(url string) (string, error) {
	var ordinaryurl string
	err := db.pool.QueryRow(context.Background(), "SELECT original_url COUNT FROM urls WHERE encoded_url = $1", url).Scan(&ordinaryurl)
	if err != nil {
		return "", err
	}
	return ordinaryurl, nil
}

func (db *Database) ExistsOrdinary(url string) bool {
	var count int
	db.pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM urls WHERE original_url = $1", url).Scan(&count)
	return count > 0
}

func (db *Database) ExistsEncoded(url string) bool {
	var count int
	db.pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM urls WHERE original_url = $1", url).Scan(&count)
	return count > 0
}

func (db *Database) InsertURL(url string, surl string) error {
	query := `INSERT INTO urls (original_url, encoded_url) VALUES ($1, $2)`
	_, err := db.pool.Exec(context.Background(), query, url, surl)
	if err != nil {
		return err
	}
	return nil
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
