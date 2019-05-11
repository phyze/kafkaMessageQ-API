package commu

import (
	"KafkaMessageQ-API/server/core/config"
	i "KafkaMessageQ-API/server/core/ideal/commu"
	"KafkaMessageQ-API/server/plugin"
	"encoding/json"
	"errors"
	"reflect"
	"sort"
	"strings"
)

//implement ideal (interface)
var _ i.CommunicationInterface = (*ProduceForm)(nil)

type ProduceForm struct {
	Await          bool                   `json:"await"`
	Message        map[string]interface{} `json:"message"`
	Topics         map[string]string      `json:"topics"`
	TimeoutProduce int                    `json:"TimeoutProduce"`
	TimeoutConsume int                    `json:"TimeoutConsume"`
	Timeout        int                    `json:"timeout"`
}

func (pf *ProduceForm) Validate() error {

	if pf.GetMessage() == "" {
		return errors.New("message is empty")
	}

	if pf.Topics == nil {
		return errors.New("topics is empty")
	}

	
	if pf.Topics[config.SendingMessageToTopic] == "" {
		return errors.New("sending is empty")
	} 

	if pf.Await == false {
		if pf.Topics[config.ReceivingMessageFromTopic] != "" {
			return errors.New("if producer is async ,thus ,receiving should be empty")
		}
	}
	
	

	return nil
}

func (pf *ProduceForm) GetFieldsName() []string {
	var sliceKeysLocal []string
	valueEF := reflect.Indirect(reflect.ValueOf(pf))
	numOfField := valueEF.Type().NumField()
	for i := 0; i < numOfField; i++ {
		sliceKeysLocal = append(sliceKeysLocal, valueEF.Type().Field(i).Name)
	}
	return sliceKeysLocal
}

//IsCompatible ...check keys of mapI compare with fields of struct
func (pf *ProduceForm) IsCompatible(mapI *map[string]interface{}) bool {

	const (
		timeoutConsume = "timeoutConsume"
		timeoutProduce = "timeoutProduce"
		timeout        = "timeout"
		aWait          = "aWait"
	)

	sliceKeysStringMapI := plugin.GetSliceStringKeyFromMap(*mapI)
	sliceKeysLocal := pf.GetFieldsName()

	//set  keys into mapI when no have keys because their must to be default
	isin, _ := plugin.Isin(timeoutProduce, sliceKeysStringMapI)
	if !isin {
		sliceKeysStringMapI = append(sliceKeysStringMapI, timeoutProduce)
	}

	isin, _ = plugin.Isin(timeoutConsume, sliceKeysStringMapI)
	if !isin {
		sliceKeysStringMapI = append(sliceKeysStringMapI, timeoutConsume)
	}

	isin, _ = plugin.Isin(timeout, sliceKeysStringMapI)
	if !isin {
		sliceKeysStringMapI = append(sliceKeysStringMapI, timeout)
	}

	isin, _ = plugin.Isin(aWait, sliceKeysStringMapI)
	if !isin {
		sliceKeysStringMapI = append(sliceKeysStringMapI, aWait)
	}

	sort.Strings(sliceKeysLocal)
	sort.Strings(sliceKeysStringMapI)
	a := strings.Join(sliceKeysLocal, " ")
	b := strings.Join(sliceKeysStringMapI, " ")
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if strings.Compare(a, b) != 0 {
		return false
	}
	return true
}

func (pf *ProduceForm) GetMessage() string {
	s, _ := json.Marshal(pf.Message)
	return string(s)
}
