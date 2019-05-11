package kafkaclient

import (
	"KafkaMessageQ-API/server/core/structs/commu"
	"context"
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	jsoniter "github.com/json-iterator/go"
)

//AwaitMessage harvest the response message from topic-res
//of produceSync after produce message
func AwaitMessage(topic string, configMap *kafka.ConfigMap, timeout int) ([]byte, error) {
	consumeError := make(chan error)
	consumeValue := make(chan []byte)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	go consume(topic, configMap, consumeValue, consumeError)

	select {
	case <-ctx.Done():
		return []byte(""), ctx.Err()
	case value := <-consumeValue:

		return value, nil
	}

}

func consume(topic string, configMap *kafka.ConfigMap, consumeValue chan []byte, consumeError chan error) {

	messageForm := commu.MessageForm{}
	emptyValue := []byte(``)
	c, err := kafka.NewConsumer(configMap)
	defer c.Close()

	c.Subscribe(topic, nil)
	if err != nil {

		consumeValue <- emptyValue
		consumeError <- err
	}

	for {

		msg, err := c.ReadMessage(-1)

		if err != nil {
			consumeValue <- msg.Value
			consumeError <- err

		}

		err = json.Unmarshal(msg.Value, &messageForm)
		if err != nil {
			consumeValue <- emptyValue
			consumeError <- err
		}

		groupID, err := configMap.Get("group.id", nil)
		if err != nil {
			consumeValue <- emptyValue
			consumeError <- err
		}
		if messageForm.ClientID == groupID.(string) {
			data, err := jsoniter.Marshal(messageForm.Message)

			if err != nil {
				consumeValue <- emptyValue
				consumeError <- err
			}
			consumeValue <- data
			consumeError <- nil
		}
	}

}
