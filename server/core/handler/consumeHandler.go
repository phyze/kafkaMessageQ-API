package handler

import (
	"errors"
	"kafkaMessageQ-API/server/core/config"
	"kafkaMessageQ-API/server/core/services"
	"kafkaMessageQ-API/server/core/structs/commu"
	"kafkaMessageQ-API/server/core/structs/logger"
	"kafkaMessageQ-API/server/plugin"
	"net/http"
	"os"
	"sync"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

func ConsumeHandle(res http.ResponseWriter, req *http.Request) {
	Infofile, _ := os.OpenFile(config.InfoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, config.PermissionlogFire)

	changedFieldTimeToUTCLog := &log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "localTime",
		},
	}

	logWriter.Formatter = changedFieldTimeToUTCLog
	logWriter.Level = log.TraceLevel
	logger := logger.Logger{}

	defer req.Body.Close()
	var consumeForm *commu.ConsumerForm
	mapBody := make(map[string]interface{})
	bodyRaw, err := plugin.ReadBody(req)

	if err != nil {

		writeLog(false, "info", &logger, Infofile, err.Error(), req)

		http.Error(res, plugin.ErrorResponseJson(err), _badRequest)
		return
	}

	writeLog(true, "info", &logger, Infofile, bodyRaw, req)

	err = jsoniter.Unmarshal(bodyRaw, &mapBody)
	if err != nil {

		writeLog(false, "info", &logger, Infofile, err.Error(), req)

		http.Error(res, plugin.ErrorResponseJson(err), _badRequest)
		return
	}

	if err = jsoniter.Unmarshal(bodyRaw, &consumeForm); err != nil {

		writeLog(false, "info", &logger, Infofile, err.Error(), req)

		http.Error(res, plugin.ErrorResponseJson(err), _serverError)
		return
	}

	if !consumeForm.IsCompatible(&mapBody) {
		err = errors.New("json incompatible")

		writeLog(false, "info", &logger, Infofile, err, req)

		http.Error(res, plugin.ErrorResponseJson(err), _badRequest)
		return
	}

	//check validate data in struct
	if err = consumeForm.Validate(); err != nil {

		writeLog(false, "info", &logger, Infofile, err.Error(), req)

		http.Error(res, plugin.ErrorResponseJson(err), _badRequest)
		return
	}

	responseTrigger := make(chan *commu.ResponseTringger)
	var wg sync.WaitGroup
	wg.Add(1)
	go services.ConsumeService(consumeForm, responseTrigger, &wg)
	response := <-responseTrigger
	data, err := jsoniter.Marshal(response)
	if err != nil {

		writeLog(false, "info", &logger, Infofile, err.Error(), req)
		http.Error(res, plugin.ErrorResponseJson(err), _serverError)
		return
	}

	writeLog(false, "info", &logger, Infofile, data, req)
	res.Write(data)
	wg.Wait()

}
