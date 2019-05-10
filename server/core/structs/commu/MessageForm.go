package commu

import (
	"encoding/json"
	"errors"
	i "kafkaMessageQ-API/server/core/ideal/commu"
	"kafkaMessageQ-API/server/plugin"
	"log"
	"reflect"
	"sort"
	"strings"
)

//implement ideal (interface)
var _ i.CommunicationInterface = (*MessageForm)(nil)

type MessageForm struct {
	Message  map[string]interface{} `json:"message"`
	ClientID string                 `json:"clientID"`
}

func (mf *MessageForm) Validate() error {

	if mf.GetMessage() == "" {
		return errors.New("message is empty")
	}
	if mf.ClientID == "" {
		return errors.New("UUID is empty")
	}
	return nil
}

func (mf *MessageForm) GetFieldsName() []string {
	var sliceKeysLocal []string
	valueEF := reflect.Indirect(reflect.ValueOf(mf))
	numOfField := valueEF.Type().NumField()
	for i := 0; i < numOfField; i++ {
		sliceKeysLocal = append(sliceKeysLocal, valueEF.Type().Field(i).Name)
	}
	return sliceKeysLocal
}

//IsCompatible ...check keys of mapI compare with fields of struct
func (mf *MessageForm) IsCompatible(mapI *map[string]interface{}) bool {

	sliceKeysStringMapI := plugin.GetSliceStringKeyFromMap(*mapI)
	sliceKeysLocal := mf.GetFieldsName()

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

func (mf *MessageForm) GetMessage() string {
	s, err := json.Marshal(mf.Message)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(s)
}
