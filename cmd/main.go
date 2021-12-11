package main

import (
	"fmt"

	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	cfg := config.Load()
	fmt.Println(logger, cfg)
}
