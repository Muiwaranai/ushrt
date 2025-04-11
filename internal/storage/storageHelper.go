package storage

import (
	"errors"
	"fmt"
	"os"
)

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
