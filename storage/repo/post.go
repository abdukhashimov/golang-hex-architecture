package repo

import "github.com/abdukhashimov/golang-hex-architecture/service/models"

type PostI interface {
	Create(payload models.Post) (string, error)
	GetUserPosts(filter models.PostFilter) (models.PostAll, error)
}
