package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/tmnhs/fginx/server/global"
	"go.uber.org/zap"
	"strings"
)

//消息队列 --消费者
type ConsumerService struct {
}

var consumer sarama.Consumer

type ConsumerCallback func(data []byte)

//初始化消费者
func (consumerService *ConsumerService) InitConsumer(hosts string) error {
	config := sarama.NewConfig()
	//config.Consumer.Return.Errors = true
	//config.Version = sarama.V0_11_0_2

	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		global.GV_LOG.Error("init kafka consumer client error:", zap.Error(err))
		return err
	}
	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		global.GV_LOG.Error("init kafka consumer error:", zap.Error(err))
		return err
	}
	global.GV_LOG.Debug("init consumer successful with  host", zap.Any("kafka", hosts))
	return nil
}

//消费消息,通过回调函数执行

func (consumerService *ConsumerService) ConsumeMessage(callBack ConsumerCallback) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		global.GV_LOG.Error("iConsumePartition error", zap.Error(err))
		return
	}
	defer partitionConsumer.Close()
	global.GV_LOG.Debug("consume topic ", zap.Any("kafka", topic))
	for {

		//select {
		//case msg := <-partitionConsumer.Messages():
		//	global.GV_LOG.Debug("kafka consume message")
		//	if callBack!=nil {
		//		callBack(msg.Value)
		//	}
		//case err := <-partitionConsumer.Errors():
		//	global.GV_LOG.Error("ConsumeMsg error:",zap.Error(err))
		//}
		msg := <-partitionConsumer.Messages()
		global.GV_LOG.Debug("kafka consume message")
		if callBack != nil {
			callBack(msg.Value)
		}
	}
}

func (consumerService *ConsumerService) Close() {
	if consumer != nil {
		_ = consumer.Close()
	}
}
