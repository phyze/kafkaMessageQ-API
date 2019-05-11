package services

import (
	"KafkaMessageQ-API/server/core/config"
	"KafkaMessageQ-API/server/core/structs/commu"
	"KafkaMessageQ-API/server/plugin/kafkaclient"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

func ConsumeService(cf *commu.ConsumerForm, responseTrigger chan *commu.ResponseTringger, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		response commu.ResponseTringger
		mf       commu.MessageForm
	)

	//check topics
	kafkaResponse, err := kafkaclient.ExistsTopic(cf.Topics, &config.ConfigConsume)

	//if kafka was error return error that receive from kafka

	//error from server
	if kafkaResponse == 0 && err == nil {

		response.Error = err.Error()
		response.StatusCode = _serverError
		responseTrigger <- &response
		return
	}

	//topic not found
	if err != nil && kafkaResponse == 0 {

		response.Error = err.Error()
		response.StatusCode = _notFound
		responseTrigger <- &response
		return
	}

	//if slice of topcs == []  by default will subscribe all topics
	if kafkaResponse == -1 {

		metadata, err := kafkaclient.GetTopics(&config.ConfigConsume)

		if err != nil {
			response.Error = err.Error()
			response.StatusCode = _serverError
			responseTrigger <- &response

			return
		}
		sliceTopics := kafkaclient.GetSliceKeyFromTopicsMetadata(metadata)
		regexp, _ := regexp.Compile(config.PatternFindStringSubmatchTopics)
		sliceTopics = regexp.FindAllString(strings.Join(sliceTopics, " "), -1)

		cf.Topics = sliceTopics

	}

	if !cf.HasAutoOffsetReset() {

		err := fmt.Sprintf(`{"error":"autoOffsetReset: %s isn't allowed"}`, cf.AutoOffsetReset)
		response.Error = errors.New(err).Error()
		response.StatusCode = _badRequest
		responseTrigger <- &response
		return
	}

	configMap := config.ConfigConsume
	configMap["group.id"] = cf.GroupID
	if cf.AutoOffsetReset != "" {
		configMap["auto.offset.reset"] = cf.AutoOffsetReset
	}
	msg, err := kafkaclient.Consume(cf.Topics, &configMap)

	if err != nil {
		response.Error = err.Error()
		response.StatusCode = _serverError
		responseTrigger <- &response
		return
	}

	if err = jsoniter.Unmarshal(msg, &mf); err != nil {
		response.Error = err.Error()
		response.StatusCode = _serverError
		responseTrigger <- &response
		return
	}

	//response
	response.Result = mf
	response.StatusCode = _ok

	responseTrigger <- &response
}
