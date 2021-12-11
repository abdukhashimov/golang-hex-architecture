package v1

import (
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/service"
	"go.uber.org/zap"
)

type handlerV1 struct {
	log     *zap.Logger
	cfg     *config.Config
	service service.ServiceI
}

type HandlerOptions struct {
	Cfg     *config.Config
	Log     *zap.Logger
	Service service.ServiceI
}

func NewHandler(opts *HandlerOptions) *handlerV1 {
	return &handlerV1{
		cfg:     opts.Cfg,
		log:     opts.Log,
		service: opts.Service,
	}
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
