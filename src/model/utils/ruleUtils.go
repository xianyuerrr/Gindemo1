package utils

import (
	"encoding/json"
	"grayRelease/src/model/tables"
)

func CreateRuleFromString(string string) *tables.Rule {
	var rule tables.Rule
	err := json.Unmarshal([]byte(string), &rule)
	if err != nil {
		return nil
	}
	return &rule
}

func CreateNewRuleFromString(string string) *tables.NewRule {
	var newRule tables.NewRule
	err := json.Unmarshal([]byte(string), &newRule)
	if err != nil {
		return nil
	}
	return &newRule
}
