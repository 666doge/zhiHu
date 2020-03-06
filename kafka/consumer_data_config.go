package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type consumerConfig struct {
	Addr string
	Topic string
	Handler func(message *sarama.ConsumerMessage)
}

var consumerList = []consumerConfig{
	consumerConfig{
		Addr: "localhost:9092",
		Topic: "zhihu_question",
		Handler: func(msg *sarama.ConsumerMessage){
			fmt.Print("got question from kafka, msg: %v", msg)
			return
		},
	},
	consumerConfig{
		Addr: "localhost:9092",
		Topic: "zhihu_answer",
		Handler: func(msg *sarama.ConsumerMessage){
			fmt.Print("got answer from kafka, msg: %v", msg)
			return
		},
	},
}