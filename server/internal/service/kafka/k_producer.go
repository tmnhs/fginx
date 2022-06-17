package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/tmnhs/fginx/server/global"
	"go.uber.org/zap"
	"strings"
)

//消息队列 --生产者
type ProducerService struct {
}

var producer sarama.AsyncProducer

var topic string

func (producerService *ProducerService) InitProducer(topicInput string, hosts string) error {
	if topicInput == "" || len(topicInput) == 0 {
		topic = global.GV_CONFIG.Kafka.Topic
	}
	topic = topicInput
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP
	//config.Producer.RequiredAcks = sarama.WaitForAll
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	//config.Producer.Return.Successes = true
	//config.Producer.Return.Errors = true
	//config.Version = sarama.V0_11_0_2
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		global.GV_LOG.Error("init kafka client error", zap.Error(err))
		return err
	}

	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		global.GV_LOG.Error("init kafka async client error", zap.Error(err))
		return err
	}
	global.GV_LOG.Debug("init producer successful with  host", zap.Any("kafka", hosts))

	return nil
}

func (producerService *ProducerService) Send(data []byte) {
	//global.GV_LOG.Debug("kafka produce message")
	b := sarama.ByteEncoder(data)

	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: b}
}

func (producerService *ProducerService) Close() {
	if producer != nil {
		_ = producer.Close()
	}
}
