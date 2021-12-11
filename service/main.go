package service

import (
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/storage"
	"go.uber.org/zap"
)

type serviceHandler struct {
	postService PostI
}

type ServiceI interface {
	Post() PostI
}

func NewServiceHandler(cfg *config.Config, log *zap.Logger, strg storage.StorageI) ServiceI {
	return &serviceHandler{
		postService: NewPostService(strg),
	}
}

func (s *serviceHandler) Post() PostI {
	return s.postService
}
