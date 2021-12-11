package service

import (
	"context"

	"github.com/abdukhashimov/golang-hex-architecture/service/models"
)

type PostI interface {
	CreatePost(context.Context, models.Post) (string, error)
	GetPostsByAuthhor(context.Context, models.PostFilter) (models.PostAll, error)
}
