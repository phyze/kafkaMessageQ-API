package kafkaclient

import (
	"KafkaMessageQ-API/server/core/config"
	"errors"
	"reflect"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//ExitsTopic if topics is exists return (1, nil),
//if not exsits return (0, nil),
//if error return (0, error)
//if error and no values in slice of topics return -1, error
//topics type that particular allowed  string , []string , map[string]string  that
//require sending and receiving key:value
func ExistsTopic(topics interface{}, kafkaConfig *kafka.ConfigMap) (int, error) {

	var dummyTopic []string

	if reflect.TypeOf(topics) == reflect.TypeOf(dummyTopic) {
		//checking size of slice of topics for protect memory addr error
		//because that memory addr error at switch
		if len(topics.([]string)) == 0 {
			return -1, errors.New("no have values in slice of topics")
		}
	}

	admin, err := kafka.NewAdminClient(kafkaConfig)
	if err != nil {
		return 0, err
	}

	switch topics := topics.(type) {

	case []string:

		// e, _ := admin.GetMetadata(&topics[0], false, 30)
		for _, topic := range topics {
			metadata, err := admin.GetMetadata(&topic, false, config.TimoutGetMetaDataFromKafkaMS)
			if err != nil {
				return 0, err
			}

			if metadata.Topics[topic].Error.Error() != "Success" {
				return 0, metadata.Topics[topic].Error
			}
		}

	case string:

		metadata, err := admin.GetMetadata(&topics, false, config.TimoutGetMetaDataFromKafkaMS)
		if err != nil {
			return 0, err
		}

		if metadata.Topics[topics].Error.Error() != "Success" {
			return 0, metadata.Topics[topics].Error
		}

	case *map[string]string:

		var topicDeref = *topics

		//if aWait is false
		sending := topicDeref[config.SendingMessageToTopic]
		metadata, err := admin.GetMetadata(&sending, false, config.TimoutGetMetaDataFromKafkaMS)
		if err != nil {
			return 0, err
		}

		if metadata.Topics[sending].Error.Error() != "Success" {
			return 0, metadata.Topics[sending].Error
		}

		//if aWait is true
		receiving := topicDeref[config.ReceivingMessageFromTopic]
		if receiving != "" {
			metadata, err = admin.GetMetadata(&receiving, false, config.TimoutGetMetaDataFromKafkaMS)
			if err != nil {
				return 0, err
			}

			if metadata.Topics[receiving].Error.Error() != "Success" {
				return 0, metadata.Topics[receiving].Error
			}
		}

	}

	return 1, nil
}

func GetTopics(kafkaConfig *kafka.ConfigMap) (map[string]kafka.TopicMetadata, error) {
	topics := make(map[string]kafka.TopicMetadata)
	admin, err := kafka.NewAdminClient(kafkaConfig)
	if err != nil {
		return topics, err
	}
	meta, err := admin.GetMetadata(nil, true, config.TimoutGetMetaDataFromKafkaMS)
	if err != nil {
		return topics, err
	}
	return meta.Topics, nil
}

func GetSliceKeyFromTopicsMetadata(m map[string]kafka.TopicMetadata) []string {
	var s []string
	for topic, _ := range m {
		s = append(s, topic)
	}
	return s
}
