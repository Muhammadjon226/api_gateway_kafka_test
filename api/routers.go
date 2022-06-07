package api

import (
	_ "github.com/Muhammadjon226/api_gateway/api/docs" // swag
	v1 "github.com/Muhammadjon226/api_gateway/api/handlers/v1"
	"github.com/Muhammadjon226/api_gateway/config"
	"github.com/Muhammadjon226/api_gateway/pkg/event"
	"github.com/Muhammadjon226/api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Config struct {
	Logger logger.Logger
	Config config.Config
	Kafka  *event.Kafka
}

// New is a constructor for gin.Engine
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cfg Config) *gin.Engine {
	if cfg.Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// this html has been copied to Dockerfile
	// r.LoadHTMLGlob("html/**/*")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	HandlerV1 := v1.New(
		cfg.Logger,
		cfg.Config,
		cfg.Kafka,
	)

	r.POST("/v1/user/create-user/", HandlerV1.CreatUser)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
