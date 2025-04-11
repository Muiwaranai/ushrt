package service

import "ushrt/internal/storage"

type Service struct {
	db *storage.Database
}

func New(db *storage.Database) *Service {
	return &Service{
		db: db,
	}
}
