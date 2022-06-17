package initialize

import (
	"fmt"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/service"
	"go.uber.org/zap"
)

var kafkaService = service.ServiceGroupApp.KafkaServiceGroup

func Kafka() {
	if global.GV_CONFIG.Kafka.Use == true {
		err := kafkaService.ConsumerService.InitConsumer(fmt.Sprintf("%s:%s", global.GV_CONFIG.Kafka.Path, global.GV_CONFIG.Kafka.Port))
		if err != nil {
			global.GV_LOG.Error("Kafka  InitConsumer err:", zap.Error(err))
			return
		}
		err = kafkaService.ProducerService.InitProducer(global.GV_CONFIG.Kafka.Topic, fmt.Sprintf("%s:%s", global.GV_CONFIG.Kafka.Path, global.GV_CONFIG.Kafka.Port))
		if err != nil {
			global.GV_LOG.Error("Kafka  InitProducer err:", zap.Error(err))
			return
		}
		global.GV_LOG.Debug("Kafka  Init successful")
	}
	//go kafkaService.ConsumeMessage()
}
