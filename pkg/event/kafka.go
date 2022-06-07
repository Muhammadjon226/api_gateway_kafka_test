package event

import (
	"github.com/Shopify/sarama"
	"github.com/Muhammadjon226/api_gateway/config"
	"github.com/Muhammadjon226/api_gateway/pkg/logger"
)

//Kafka ...
type Kafka struct {
	log          logger.Logger
	cfg          config.Config
	publishers   map[string]*Publisher
	saramaConfig *sarama.Config
}

//NewKafka ...
func NewKafka(cfg config.Config, log logger.Logger) (*Kafka, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0

	kafka := &Kafka{
		log:          log,
		cfg:          cfg,
		publishers:   make(map[string]*Publisher),
		saramaConfig: saramaConfig,
	}

	kafka.RegisterPublishers()

	return kafka, nil
}

//RegisterPublishers ...
func (k *Kafka) RegisterPublishers() {

	// user service
	userRoute := "v1.user"
	k.AddPublisher(userRoute + ".created")
	k.AddPublisher(userRoute + ".updated")
	// k.AddPublisher(userRoute + ".list")
	// k.AddPublisher(userRoute + ".get")
	k.AddPublisher(userRoute + ".deleted")

	
}
