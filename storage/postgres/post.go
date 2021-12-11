package postgres

import (
	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/abdukhashimov/golang-hex-architecture/storage/repo"
	"github.com/jmoiron/sqlx"
)

type posts struct {
	db *sqlx.DB
}

func NewPostStorage(db *sqlx.DB) repo.PostI {
	return &posts{db: db}
}

func (p *posts) Create(payload models.Post) (string, error) {
	return "", nil
}

func (p *posts) GetUserPosts(filter models.PostFilter) (models.PostAll, error) {
	return models.PostAll{}, nil
}

func (p *posts) GetPosts(page, limit int) (models.PostAll, error) {
	return models.PostAll{}, nil
}
