// +build !prod

package testing

import (
	"KafkaMessageQ-API/server/core/config"
	"KafkaMessageQ-API/server/plugin/kafkaclient"
	"log"
	"testing"
)

func TestGetTopiceMetadata(t *testing.T) {
	topics, err := kafkaclient.GetTopics(&config.AsyncConfigProduce)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Fatal(kafkaclient.GetSliceKeyFromTopicsMetadata(topics))

}
