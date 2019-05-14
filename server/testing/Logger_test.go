// +build !prod

package testing

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func TestWritelog(t *testing.T) {
	logger := log.New()
	logger.Formatter = new(log.JSONFormatter)
	logger.Level = log.TraceLevel
	logger.Out = os.Stdout

	file, err := os.OpenFile("/Users/a./go/src/KafkaMessageQ-API/logs/info/info.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}


	logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
