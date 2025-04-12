package service

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"ushrt/internal/storage"

	"github.com/jackc/pgx/v5"
)

type Service struct {
	db *storage.Database
}

func New(db *storage.Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Hash2URL(short string) (string, error) {
	url, err := s.db.ByShort(short)
	if err == pgx.ErrNoRows {
		return "", fmt.Errorf("404 page not found")
	} else if err != nil {
		return "", err
	}

	url, err = SanitizeURL(url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *Service) URL2Hash(url string) (string, error) {
	url, err := SanitizeURL(url)
	if err != nil {
		return "", err
	}

	short, err := s.db.ByURL(url)
	if err == pgx.ErrNoRows {
		return s.createShort(url)
	}

	return short, err
}

func (s Service) createShort(url string) (string, error) {
	errCount := 0
	for {
		short := encodeShort()
		if yes, err := s.checkShort(short); !yes {
			if err != nil {
				return "", err
			}

			err = s.db.Insert(url, short)

			if err != nil {
				return "", err
			}
			return short, nil
		}
		errCount++
		if errCount == 10 {
			return "", errors.New("Error with encoding url")
		}
	}
}

func (s *Service) checkShort(short string) (bool, error) {
	_, err := s.db.ByShort(short)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	return true, err
}

func encodeShort() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hash := []byte{'/', 'r', '/'}
	for i := 0; i < 8; i++ {
		hash = append(hash, charset[rand.Intn(len(charset)-1)])
	}
	return string(hash)
}

func SanitizeURL(u string) (string, error) {
	if u == "" {
		return "", errors.New("Bad url format")
	}
	url, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	url.User = nil
	url.Scheme = "http"

	splitted := strings.Split(url.String(), "://")
	if len(splitted) != 2 {
		return "", fmt.Errorf("invalid URL")
	}

	return splitted[1], nil
}
