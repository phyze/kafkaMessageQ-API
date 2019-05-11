// +build !dev

package services

import (
	"errors"
	"fmt"
	"KafkaMessageQ-API/server/core/config"
	"KafkaMessageQ-API/server/core/structs/commu"
	"KafkaMessageQ-API/server/plugin"
	"KafkaMessageQ-API/server/plugin/kafkaclient"
	"KafkaMessageQ-API/server/plugin/uuid"
	"regexp"
	"strings"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	jsoniter "github.com/json-iterator/go"
)

func ProduceService(pf *commu.ProduceForm, rt chan *commu.ResponseTringger) {
	var mf commu.MessageForm
	var response commu.ResponseTringger
	var wg sync.WaitGroup

	//Check the topic is exists in Kafka or not.
	kafkaResponse, err := kafkaclient.ExistsTopic(&pf.Topics, &config.AsyncConfigProduce)
	//if kafka was error return error that receive from kafka

	if err != nil && kafkaResponse == 0 {
		response.Error = err.Error()
		response.StatusCode = _serverError
		rt <- &response
		return
	} else if err != nil && kafkaResponse == -1 {
		response.Error = err.Error()
		response.StatusCode = _badRequest
		rt <- &response
		return
	} else if kafkaResponse == 0 && err == nil {
		err = errors.New("topics not found")
		response.Error = err.Error()
		response.StatusCode = _badRequest
		rt <- &response
		return
	}

	if pf.Await {
		wg.Add(1)
		errChan := make(chan error)
		go producer(pf.Topics[config.SendingMessageToTopic], true, pf, &wg, errChan)

		err = <-errChan

		if err != nil {
			response.Error = err.Error()
			response.StatusCode = _serverError
			rt <- &response
			return
		}

		configMap := config.ConfigConsume
		configMap["group.id"] = clientID.String()

		//คำนวลหาว่า topic ขา sending มี -req เข้ามาไหม และ receving ค่าไหมถ้าไม่
		//ให้ดึง topice name จาก sending จากนั้นเติม -res และส่งไปที่ func AwaitMessage
		receivingTopic := pf.Topics[config.ReceivingMessageFromTopic]
		sendingTopic := pf.Topics[config.SendingMessageToTopic]
		if receivingTopic == "" && sendingTopic != "" {
			if TopicHasEndFix(pf.Topics[config.SendingMessageToTopic]) {
				topicNameForConsumer := MidgTopicNames(sendingTopic)
				if strings.Contains(FindTopicBody(sendingTopic), topicNameForConsumer) {
					receivingTopic = topicNameForConsumer + "-" + config.EndFixProducerConsume
				}
			}
		}

		//block until receive the message
		data, err := kafkaclient.AwaitMessage(receivingTopic, &configMap, pf.TimeoutConsume)

		if err != nil {
			response.Error = err.Error()
			response.StatusCode = _serverError
			rt <- &response
			return
		}

		if err = jsoniter.Unmarshal(data, &mf.Message); err != nil {
			response.Error = err.Error()
			response.StatusCode = _serverError
			rt <- &response
			return
		}

		//response data to client
		mf.ClientID = clientID.String()
		response.Result = mf
		response.StatusCode = _ok

		rt <- &response
		return
	} else {
		//not await
		wg.Add(1)
		// go produceAsync(topic, &pf, &wg)
		errChan := make(chan error)

		go producer(pf.Topics[config.SendingMessageToTopic], false, pf, &wg, errChan)
		err = <-errChan

		if err != nil {
			response.Error = err.Error()
			response.StatusCode = _serverError
			rt <- &response
			return
		}

		//return no content
		response.StatusCode = _noContent
		rt <- &response
		return
	}

	wg.Wait()
}

func producer(topic string, identify bool, pf *commu.ProduceForm, wg *sync.WaitGroup, errChan chan error) {
	templatePayload := `{"message":{{.message}},"clientID":"{{.clientID}}"}`
	defer wg.Done()

	var id uuid.UUID

	//gen ClientID
	if identify {
		if id, err = uuid.NewUUID4(); err != nil {
			errChan <- err
			return
		}
	}

	if err != nil {
		errChan <- err //buffer error  to chan error if uuid had error
		return
	}

	//check clientID if exists wil not gen new.
	//this condition use with consume need to produce message to
	//producer that using sync style
	if pf.Message["clientID"] != nil {
		idString := pf.Message["clientID"].(string)
		idObj, err := uuid.Parse(idString)
		if err != nil {
			errChan <- err
			return
		}
		clientID = &idObj
		delete(pf.Message, "clientID")
	} else if pf.Message["clientID"] == nil {
		clientID = &id
	}

	payloadMap := map[string]interface{}{
		"message":  pf.GetMessage(),
		"clientID": clientID.String(),
	}

	payload := plugin.StringFormat(templatePayload, payloadMap)

	//set config for produce message
	syncConfigMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(payload),
	}

	err = kafkaclient.Produce(&config.SyncConfigProduce, syncConfigMessage, pf.TimeoutProduce)

	if err != nil {
		errChan <- err
		return
	}
	errChan <- nil
	return
}

//TopicHasEndFix เช็คว่ามี -req หรือไม่ ถ้าไม่มี return false
func TopicHasEndFix(s string) bool {
	shatters := strings.Split(s, "-")
	if shatters[len(shatters)-1] == config.EndfixProducer {
		return true
	}
	return false
}

//FindTopicBody ดึงค่า topic name จาก parameter และเอาแค่ส่วน body ไม่เอา -req
func FindTopicBody(s string) string {
	regex, _ := regexp.Compile(fmt.Sprintf("(?:([a-zA-Z0-9-]+))(?:(-%s))", config.EndfixProducer))
	slice := regex.FindAllStringSubmatch(s, -1)
	return strings.Join(slice[0][:], " ")
}

//MidTopicNames ดึงค่ากลางของ topic name ที่ได้จากการแตก topic name ให้เป็นหลายค่า
//
//::Pattern:: (?:([a-zA-Z0-9-]+))(?:(-req))
//
//::example::  abc-req => abc-req, abc, -req โดยหยิปค่าตรงกลางเสมอ
func MidgTopicNames(s string) string {

	regex, _ := regexp.Compile(fmt.Sprintf("(?:([a-zA-Z0-9-]+))(?:(-%s))", config.EndfixProducer))
	slice := regex.FindAllStringSubmatch(s, -1)
	return slice[0][1]
}
