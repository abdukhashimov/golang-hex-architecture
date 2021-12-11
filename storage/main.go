package storage

import (
	"github.com/abdukhashimov/golang-hex-architecture/storage/postgres"
	"github.com/abdukhashimov/golang-hex-architecture/storage/repo"
	"github.com/jmoiron/sqlx"
)

type storage struct {
	postStorage repo.PostI
}

type StorageI interface {
	Post() repo.PostI
}

func NewStorage(db *sqlx.DB) StorageI {
	return &storage{
		postStorage: postgres.NewPostStorage(db),
	}
}

func (s *storage) Post() repo.PostI {
	return s.postStorage
}
