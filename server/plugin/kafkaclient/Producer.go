package kafkaclient

import (
	"kafkaMessageQ-API/server/core/config"
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//Produce is produce message to kafka broker
func Produce(configMap *kafka.ConfigMap, configMessage *kafka.Message, timeout int) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()
	errEvent := make(chan error)
	errProduce := make(chan error)
	producer, err := kafka.NewProducer(configMap)
	
	if err != nil {
		return err
	}

	//listening result of producer to message whether sending succeed or not
	//if succeeded ,which TopicPartition.Error will return nil
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				errEvent <- ev.TopicPartition.Error
			}
		}
	}()

	go func() {
		defer producer.Close()
		if err = producer.Produce(configMessage, nil); err != nil {
			errProduce <- err
		}
		producer.Flush(config.FlushProduceTime)
	}()

	//checking timeout producing message
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errEvent:
		//in this case wheather produce succeeded or produce failed
		//will not timeout
		if err != nil {
			return err
		}

	}

	return nil
}
