package kafka

import (
	"zhiHu/logger"
	"strings"
	"encoding/json"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	producer sarama.SyncProducer
	wg sync.WaitGroup
)

func InitProducer(addr []string)(err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	producer, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		return
	}
	logger.Info("kafka is initialized at %#v", addr)
	return
}

func SendMessage(topic string, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		logger.Error("send to kafka : encode value failed, err: %v", err)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logger.Error("message send to kafka Failed, err: %v", err)
	} else {
		logger.Info("message send to kafka Success, pid: %v, offset: %v, data: %v", partition, offset, data)
	}
	return
}

func InitConsumers() {
	for _, config := range consumerList {
		go createConsumer(
			config.Addr,
			config.Topic,
			config.Handler,
		)
	}
}

func createConsumer(addrs, topic string, consume func(message *sarama.ConsumerMessage)) (err error) {
	consumer, err := sarama.NewConsumer(strings.Split(addrs, ","), nil)
	if err != nil {
		logger.Error("Failed to start consumer: %s", err)
		return
	}
	defer consumer.Close()

	partitionIds, err := consumer.Partitions(topic)
	if err != nil {
		logger.Error("Failed to get the list of partitions, err: %v ", err)
		return
	}

	for pid := range partitionIds {
		partitionConsumer, err := consumer.ConsumePartition(topic, int32(pid), sarama.OffsetNewest)
		if err != nil {
			logger.Error("Failed to start consumer for partition %d: %s\n", pid, err)
		}
		wg.Add(1)
		defer partitionConsumer.Close()
		go func (pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				consume(msg)
			}
			wg.Done()
		}(partitionConsumer)
	}
	wg.Wait()
	return
}