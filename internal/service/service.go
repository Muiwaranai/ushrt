package service

import (
	"fmt"
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

func (s *Service) EncodeAndSaveURL(url string) string {

	encodedUrl := localEncode(url)

	// Save to storage

	return encodedUrl
}

func (s *Service) GetOrdinary(url string) string {
	return ""
}

func localEncode(url string) string {
	fmt.Println("qewqwe", url)
	result := make([]byte, 11, 11)
	result[0], result[1], result[2] = '/', 'r', '/'
	for i := 3; i < 11; i++ {
		result[i] = byte(71 + i)
	}
	return string(result)
}
