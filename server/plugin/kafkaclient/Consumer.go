package kafkaclient

import (
	"KafkaMessageQ-API/server/core/structs/commu"

	jsoniter "github.com/json-iterator/go"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(topics []string, configMap *kafka.ConfigMap) ([]byte, error) {

	messageForm := commu.MessageForm{}
	emptyValue := []byte(``)
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		return emptyValue, err
	}

	defer c.Close()

	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			return msg.Value, err
		}

		if err = jsoniter.Unmarshal(msg.Value, &messageForm); err != nil {
			return emptyValue, err
		}

		data, err := jsoniter.Marshal(messageForm)
		if err != nil {
			return emptyValue, err
		}

		return data, err
	}
}
