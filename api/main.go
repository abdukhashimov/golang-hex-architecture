package api

import (
	v1 "github.com/abdukhashimov/golang-hex-architecture/api/handlers/v1"
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouterOptions struct {
	Cfg *config.Config
	Log *zap.Logger
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.NewHandler(&v1.HandlerOptions{
		Cfg: opt.Cfg,
		Log: opt.Log,
	})

	apiV1 := router.Group("/v1")
	{
		apiV1.GET("/posts", handlerV1.GetUserPosts)
	}

	return router
}
