package helper

import (
	"encoding/json"
)

// TransformValue func
func TransformValue(input interface{}, output interface{}) error {
	jsonObj, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonObj, output)
	return err
}
