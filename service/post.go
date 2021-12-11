package service

import (
	"context"

	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/abdukhashimov/golang-hex-architecture/storage"
)

type PostI interface {
	CreatePost(context.Context, models.Post) (string, error)
	GetPostsByAuthhor(context.Context, models.PostFilter) (models.PostAll, error)
}

type post struct {
	db storage.StorageI
}

func NewPostService(db storage.StorageI) PostI {
	return &post{
		db: db,
	}
}

func (p *post) CreatePost(ctx context.Context, req models.Post) (string, error) {
	res, err := p.db.Post().Create(req)

	if err != nil {
		return "", err
	}

	return res, nil
}

func (p *post) GetPostsByAuthhor(ctx context.Context, filter models.PostFilter) (models.PostAll, error) {
	var (
		res = models.PostAll{}
		err error
	)

	if filter.Author == "" {
		res, err = p.db.Post().GetPosts(filter.Offset, filter.Limit)
	} else {
		res, err = p.db.Post().GetUserPosts(filter)
	}

	return res, err
}
