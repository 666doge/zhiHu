package kafka

import (
	// "fmt"
	"zhiHu/es"
	"zhiHu/model"
	"zhiHu/logger"
	"encoding/json"

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
			var question model.Question
			err := json.Unmarshal(msg.Value, &question)
			if err != nil {
				logger.Error("kafka: Unmarshal failed, err: %v", err)
				return
			}
			es.InsertDoc("zhihu_question", msg.Value, question.QuestionId, question)
			return
		},
	},
	consumerConfig{
		Addr: "localhost:9092",
		Topic: "zhihu_answer",
		Handler: func(msg *sarama.ConsumerMessage){
			var answer model.Answer
			err := json.Unmarshal(msg.Value, &answer)

			if err != nil {
				logger.Error("kafka: Unmarshal answer failed, err: %v", err)
				return
			}
			es.InsertDoc("zhihu_answer", msg.Value, answer.AnswerId, answer)
			return
		},
	},
}