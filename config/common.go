package config

import "encoding/json"

func ToJSONStruct(value interface{}, result any) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return err
	}
	return nil
}
