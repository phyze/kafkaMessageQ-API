package handler

import (
	"fmt"
	"KafkaMessageQ-API/server/core/config"
	"KafkaMessageQ-API/server/core/structs/logger"
	"KafkaMessageQ-API/server/plugin"
	"net/http"
	"os"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

const (
	_badRequest  int = 400
	_ok          int = 200
	_notFound    int = 404
	_serverError int = 500
	_noContent   int = 204
)

var (
	logWriter = log.New()
)

func Index(res http.ResponseWriter, req *http.Request) {
	i := make(map[string]interface{})
	i["Application"] = "Aggregate messaging controller API"
	i["license"] = "Gram kittisak P."
	data, _ := jsoniter.Marshal(i)
	res.Write(data)
}

//writeLog
//:: field ::
//
//communicateSC: is communication of client and server  that is request and response
//if true will use "Client's  request" ,false will use "Server's response"
//
//typeError: is what type that you need to write log
//
//logger: struct of logger.Logger
//
//typeFile: that type of log file such as errfile , infofile and debugfile
//
//data :infomation of system that you need to write
//
//req:   http.Request
func writeLog(comunicateSC bool, typeError string, logger *logger.Logger, typeFile *os.File, data interface{}, req *http.Request) {
	resOrReq := "Server's Response"
	if comunicateSC {
		resOrReq = "Client's Request"
	}
	logWriter.Out = typeFile
	logger.Data = fmt.Sprintf("%s %s %s %s %s %s", resOrReq, req.UserAgent(), req.Method, req.URL, req.RemoteAddr, data)
	logger.LogType = typeError
	logger.TimeStamp = plugin.GetDateTime(config.FormatTimeLogPattern)
	logWriter.WithFields(logger.ToMap()).Println()
}
