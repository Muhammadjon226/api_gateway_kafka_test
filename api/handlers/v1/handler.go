package v1

import (
	"github.com/Muhammadjon226/api_gateway/config"
	"github.com/Muhammadjon226/api_gateway/pkg/event"
	"github.com/Muhammadjon226/api_gateway/pkg/logger"
)

//HandlerV1 ...
type HandlerV1 struct {
	log   logger.Logger
	cfg   config.Config
	kafka *event.Kafka
}

// New ...
func New(log logger.Logger, cfg config.Config, kafka *event.Kafka) *HandlerV1 {

	return &HandlerV1{
		cfg:   cfg,
		log:   log,
		kafka: kafka,
	}
}
