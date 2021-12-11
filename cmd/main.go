package main

import (
	"github.com/abdukhashimov/golang-hex-architecture/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("Hello world", zap.String("hello", "hello"))
}
