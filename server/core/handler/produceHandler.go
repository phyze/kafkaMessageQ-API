package handler

import (
	"AMCO/server/core/config"
	"AMCO/server/core/services"
	"AMCO/server/core/structs/commu"
	"AMCO/server/core/structs/logger"
	"AMCO/server/plugin"
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

func ProduceHandle(res http.ResponseWriter, req *http.Request) {

	Infofile, _ := os.OpenFile(config.InfoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, config.PermissionlogFire)

	changedFieldTimeToUTCLog := &log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "localTime",
		},
	}

	logWriter.Formatter = changedFieldTimeToUTCLog
	logWriter.Level = log.TraceLevel
	logger := logger.Logger{}

	addValuesFromServerConfigToLogger(&logger)

	defer req.Body.Close()
	var produceForm *commu.ProduceForm
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

	if err = jsoniter.Unmarshal(bodyRaw, &produceForm); err != nil {

		writeLog(false, "info", &logger, Infofile, err.Error(), req)

		io.WriteString(res, plugin.ErrorResponseJson(err))
		return
	}

	if !produceForm.IsCompatible(&mapBody) {
		err = errors.New("json incompatible")

		writeLog(false, "info", &logger, Infofile, err.Error(), req)

		http.Error(res, plugin.ErrorResponseJson(err), _badRequest)
		return
	}

	//check validate data in struct
	if err = produceForm.Validate(); err != nil {

		writeLog(false, "info", &logger, Infofile, err.Error(), req)

		http.Error(res, plugin.ErrorResponseJson(err), _badRequest)
		return
	}

	//if timeout not specified from client it will init timeout by default
	if produceForm.Timeout == 0 {
		produceForm.Timeout = config.Timeout
	}

	if produceForm.TimeoutProduce == 0 {
		produceForm.TimeoutProduce = config.TimeoutProduceSec
	}

	if produceForm.TimeoutConsume == 0 {
		produceForm.TimeoutConsume = config.TimeoutConsumeOfProduceSyncSec
	}

	responseTrigger := make(chan *commu.ResponseTringger)
	timeoutContext, timeoutCancel := context.WithTimeout(context.Background(),
		time.Second*time.Duration(produceForm.Timeout))
	defer timeoutCancel()

	go services.ProduceService(produceForm, responseTrigger)

	select {

	case <-timeoutContext.Done():
		err := []byte(`{"data":"","statusCode":408,"error":"request timeout"}`)
		writeLog(false, "info", &logger, Infofile, err, req)
		res.Write(err)

	case response := <-responseTrigger:

		writeLog(false, "info", &logger, Infofile, response.ToJson(), req)

		res.Write(response.ToJson())
	}

}
