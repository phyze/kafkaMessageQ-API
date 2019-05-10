package plugin

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func CompareMapKeyAndFieldName(s1 []string, s2 []string) bool {
	sort.Strings(s1)
	sort.Strings(s2)
	a := strings.Join(s1, " ")
	b := strings.Join(s2, " ")
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if strings.Compare(a, b) != 0 {
		return false
	}
	return true
}

//Isin finding target in collection 1 dimension
func Isin(target, collection interface{}) (bool, error) {

	switch t := collection.(type) {
	case []string:
		if err := getTypeForIsin(target); err != nil {
			return false, err
		}
		return findInStringArray(target, t)
	case []int:
		fmt.Println(t)
	}
	return false, nil
}

func findInStringArray(target interface{}, t []string) (bool, error) {

	for _, v := range t {
		if target == v {
			return true, nil
		}
	}
	return false, nil
}

func getTypeForIsin(target interface{}) error {
	if reflect.TypeOf(target) != reflect.TypeOf("") {
		return errors.New("type of target mismatch with collection")
	}
	return nil
}



