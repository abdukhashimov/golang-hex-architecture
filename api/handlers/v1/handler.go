package v1

import (
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"go.uber.org/zap"
)

type handlerV1 struct {
	log *zap.Logger
	cfg *config.Config
}

type HandlerOptions struct {
	Cfg *config.Config
	Log *zap.Logger
}

func NewHandler(opts *HandlerOptions) *handlerV1 {
	return &handlerV1{
		cfg: opts.Cfg,
		log: opts.Log,
	}
}
