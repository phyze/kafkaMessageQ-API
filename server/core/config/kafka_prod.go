// +build prod

package config

import "github.com/confluentinc/confluent-kafka-go/kafka"

// BrokerList kafka servers
const BrokerList = "kaf1:9092,kaf2:9093"

// AutoOffsetResetDefault auto.offset.reset for consumer
const AutoOffsetResetDefault = "latest"

//AsyncConfigProduce ..
var AsyncConfigProduce = kafka.ConfigMap{
	"bootstrap.servers": BrokerList,
}

//ConfigConsume onfigMap for Consume
var ConfigConsume = kafka.ConfigMap{
	"bootstrap.servers": BrokerList,
	"auto.offset.reset": AutoOffsetResetDefault,
}

//SyncConfigProduce configMap for produce sync style
var SyncConfigProduce = kafka.ConfigMap{
	"bootstrap.servers": BrokerList,
}

//FlushProduceTime flushing message to kafka brokers
const FlushProduceTime = 15 * 1000

//TimeoutConsumeOfProduceSyncSec ..
const TimeoutConsumeOfProduceSyncSec = 30

//ProducingTimeoutSec ..
const TimeoutProduceSec = 10

//TimoutGetMetaDataFromKafkaMS ..
const TimoutGetMetaDataFromKafkaMS = 1000

//============= don't touch ====================

//PatternFindStringSubmatchTopics must find topics that like pattern   below
//
//(any characters) or (any characters)-(any characters) or (any characters any digits)
//
//(any characters any digits)-(any characters any digits)
//
//Example. abc123, abc-abc123, abc, abc123-abc, abc123-abv324, abc123-123abc-1a2b3c
const PatternFindStringSubmatchTopics = `(\A[a-zA-Z]+[0-9]*|( +[a-zA-Z]+[0-9]*[a-zA-Z0-9-]*))`

//SendingMessageToTopic name of key of topics
//which use for aWait that allow  be true and  false
const SendingMessageToTopic = "sending"

//ReceivingMessageFromTopic name of key of topics
//which using for aWait  but  aWait is only true
const ReceivingMessageFromTopic = "receiving"

//==============================================
