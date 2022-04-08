package storage

import (
	"ewallet/storage/postgres"
	"ewallet/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Ewallet() repo.Repo
}

type storage struct {
	db      *sqlx.DB
	storage *postgres.Database
}

func NewStorage(db *sqlx.DB) *storage {
	return &storage{
		db:      db,
		storage: postgres.NewDatabase(db),
	}
}

func (s *storage) Storage() *postgres.Database {
	return s.storage
}
