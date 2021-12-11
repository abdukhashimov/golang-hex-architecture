package service

import (
	"context"

	"github.com/abdukhashimov/golang-hex-architecture/service/models"
)

type PostI interface {
	CreatePost(context.Context, models.Post) (string, error)
	GetPostsByAuthhor(context.Context, models.PostFilter) (models.PostAll, error)
}

type post struct {
}

func NewPostService() PostI {
	return &post{}
}

func (p *post) CreatePost(context.Context, models.Post) (string, error) {
	return "", nil
}

func (p *post) GetPostsByAuthhor(context.Context, models.PostFilter) (models.PostAll, error) {
	return models.PostAll{}, nil
}
