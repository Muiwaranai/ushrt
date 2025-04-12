package service

import (
	"errors"
	"math/rand"
	"strings"
	"ushrt/internal/storage"
)

type Service struct {
	db *storage.Database
}

func New(db *storage.Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) EncodeAndSaveURL(url string) (string, error) {
	if !checkValidUrl(url) {
		return "", errors.New("Invalid url")
	}

	if s.db.ExistsOrdinary(url) {
		alreadyEncoded, err := s.db.GetEncoded(url)
		if err != nil {
			return "", errors.New("Problem with geting encoded url")
		}
		return alreadyEncoded, nil
	}

	errorCounter := 0

	encodedUrl := localEncode()
	for s.db.ExistsEncoded(encodedUrl) {
		encodedUrl = localEncode()
		errorCounter++

		if errorCounter == 10 {
			return "", errors.New("Problem with generating url")
		}
	}

	err := s.db.InsertURL(url, encodedUrl)
	if err != nil {
		return "", err
	}

	return encodedUrl, nil
}

func (s *Service) GetOrdinary(url string) (string, error) {
	ordinary, err := s.db.GetOrdinary(url)
	if err != nil {
		return "", err
	}
	return ordinary, nil
}

func localEncode() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, 11, 11)
	result[0], result[1], result[2] = '/', 'r', '/'
	for i := 3; i < 11; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func checkValidUrl(url string) bool {
	if url == "" || len(url) < 11 {
		return false
	}

	splitted := strings.Split(url, "://")
	if len(splitted) != 2 {
		return false
	}

	return true
}
