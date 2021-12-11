package service

import (
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"go.uber.org/zap"
)

type serviceHandler struct {
	postService PostI
}

type ServiceI interface {
	Post() PostI
}

func NewServiceHandler(cfg *config.Config, log *zap.Logger) ServiceI {
	return &serviceHandler{}
}

func (s *serviceHandler) Post() PostI {
	return s.postService
}
