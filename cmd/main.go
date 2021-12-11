package main

import (
	"github.com/abdukhashimov/golang-hex-architecture/api"
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	cfg := config.Load()
	server := api.New(&api.RouterOptions{
		Cfg: &cfg,
		Log: logger,
	})

	server.Run(cfg.HTTPPort)
}
