package utils

import "encoding/json"

func InterfaceFromBytes(b []byte) (interface{}, error) {
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}
