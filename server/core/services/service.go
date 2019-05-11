package services

import (
	"KafkaMessageQ-API/server/core/structs/logger"
	"KafkaMessageQ-API/server/plugin/uuid"
	"log"
)

const (
	_badRequest          int = 400
	_unprocessableEntity int = 422
	_ok                  int = 200
	_notFound            int = 404
	_serverError         int = 500
	_noContent           int = 204
	_noImplimented       int = 501
	_badGateWay          int = 502
)

var (
	reqBody        []byte
	err            error
	clientID       *uuid.UUID
	instanceLogger log.Logger
	systemLog      logger.Logger
)
