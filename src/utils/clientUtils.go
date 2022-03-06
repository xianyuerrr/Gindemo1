package utils

import (
	"encoding/json"
	"grayRelease/src/model"
)

func CreateClientFromString(string string) *model.Client {
	var client model.Client
	err := json.Unmarshal([]byte(string), &client)
	if err != nil {
		return nil
	}
	return &client
}
