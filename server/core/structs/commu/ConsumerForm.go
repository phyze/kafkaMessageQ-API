package commu

import (
	i "kafkaMessageQ-API/server/core/ideal/commu"
	"kafkaMessageQ-API/server/plugin"
	"errors"
	"reflect"
	"sort"
	"strings"
)

type ConsumerForm struct {
	GroupID         string   `json:"groupID"`
	Topics          []string `json:"topics"`
	AutoOffsetReset string   `json:"autoOffsetReset"`
}

var _ i.CommunicationInterface = (*ConsumerForm)(nil)

func (cf *ConsumerForm) Validate() error {
	if cf.GroupID == "" || cf.GroupID == " " {
		return errors.New("groupID is empty")
	}

	if cf.Topics == nil {
		return errors.New("topics is empty")
	}

	return nil
}

func (cf *ConsumerForm) GetFieldsName() []string {
	var sliceKeysLocal []string
	valueCF := reflect.Indirect(reflect.ValueOf(cf))
	numOfField := valueCF.Type().NumField()
	for i := 0; i < numOfField; i++ {
		sliceKeysLocal = append(sliceKeysLocal, valueCF.Type().Field(i).Name)
	}
	return sliceKeysLocal
}

//IsCompatible ..
func (cf *ConsumerForm) IsCompatible(mapI *map[string]interface{}) bool {

	const autoOffsetReset = "autoOffsetReset"

	sliceKeysStringMapI := plugin.GetSliceStringKeyFromMap(*mapI)
	sliceKeysLocal := cf.GetFieldsName()

	//set  keys into mapI when no have keys because their must to be default
	isin, _ := plugin.Isin(autoOffsetReset, sliceKeysStringMapI)
	if !isin {
		sliceKeysStringMapI = append(sliceKeysStringMapI, autoOffsetReset)
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

//HasAutoOffsetReset  if value of AutoOffsetReset
//isn't "latest" or "earliest" return false
func (cf *ConsumerForm) HasAutoOffsetReset() bool {

	if strings.ToLower(cf.AutoOffsetReset) == "latest" || strings.ToLower(cf.AutoOffsetReset) == "earliest" {
		return true
	} else if (cf.AutoOffsetReset) == "" {
		//default condition
		return true
	}
	return false
}
