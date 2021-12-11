package main

import (
	"github.com/abdukhashimov/golang-hex-architecture/api"
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/pkg/logger"
	"github.com/abdukhashimov/golang-hex-architecture/service"
)

func main() {
	logger := logger.NewLogger()
	cfg := config.Load()
	services := service.NewServiceHandler(&cfg, logger)

	server := api.New(&api.RouterOptions{
		Cfg:     &cfg,
		Log:     logger,
		Service: &services,
	})

	server.Run(cfg.HTTPPort)
}
