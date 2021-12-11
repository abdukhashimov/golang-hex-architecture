package api

import (
	_ "github.com/abdukhashimov/golang-hex-architecture/api/docs"
	v1 "github.com/abdukhashimov/golang-hex-architecture/api/handlers/v1"
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type RouterOptions struct {
	Cfg     *config.Config
	Log     *zap.Logger
	Service service.ServiceI
}

//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.NewHandler(&v1.HandlerOptions{
		Cfg:     opt.Cfg,
		Log:     opt.Log,
		Service: opt.Service,
	})

	apiV1 := router.Group("/v1")
	{
		apiV1.GET("/posts", handlerV1.GetUserPosts)
	}
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
