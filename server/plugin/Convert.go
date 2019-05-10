package plugin

import (
	"reflect"
)

func GetFieldsFromStruct(t *struct{}) *[]string {
	var s []string
	e := reflect.ValueOf(&t).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		s = append(s, varName)
	}
	return &s
}

//MapKeyToStringKey is convert map<k> to string,
//Ex. map[name age]	-> "name age" ,
func GetStringKeyFromMap(mapI map[string]interface{}) string {
	keys := ""
	i := 0
	for key := range mapI {
		if i == len(mapI)-1 {
			keys += key
		} else {
			keys += key + " "
			i++
		}
	}
	return keys
}

func GetSliceStringKeyFromMap(mapI map[string]interface{}) []string {
	sliceKeys := make([]string, 0, len(mapI))
	for key := range mapI {
		sliceKeys = append(sliceKeys, key)
	}
	return sliceKeys
}
