package main

import (
	"go.uber.org/zap"
)

func main() {

	logger.Info("Hello world", zap.String("hello", "hello"))
}
