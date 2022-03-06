package utils

import (
	"encoding/json"
	"grayRelease/src/model/tables"
)

func CreateClientFromString(string string) *tables.Client {
	var client tables.Client
	err := json.Unmarshal([]byte(string), &client)
	if err != nil {
		return nil
	}
	return &client
}
