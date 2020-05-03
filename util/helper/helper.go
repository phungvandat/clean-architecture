package helper

import (
	"encoding/json"
	"errors"
	"reflect"
)

// TransformValue function
func TransformValue(input interface{}, output interface{}) error {
	jsonObj, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonObj, output)
	return err
}

// ConvertTypeArrayToInterfaceArray function convert a []T to an []interface{}
func ConvertTypeArrayToInterfaceArray(input interface{}) ([]interface{}, error) {
	if reflect.TypeOf(input).Kind() != reflect.Slice &&
		reflect.TypeOf(input).Kind() != reflect.Array {
		return nil, errors.New("Input must is array kind")
	}

	inputVal := reflect.ValueOf(input)
	result := make([]interface{}, inputVal.Len())
	for i := 0; i < inputVal.Len(); i++ {
		result[i] = inputVal.Index(i).Interface()
	}
	return result, nil
}
