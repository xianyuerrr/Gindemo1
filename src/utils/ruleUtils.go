package utils

import (
	"encoding/json"
	"grayRelease/src/model"
)

func CreateRuleFromString(string string) *model.Rule {
	var rule model.Rule
	err := json.Unmarshal([]byte(string), &rule)
	if err != nil {
		return nil
	}
	return &rule
}

func CreateNewRuleFromString(string string) *model.NewRule {
	var newRule model.NewRule
	err := json.Unmarshal([]byte(string), &newRule)
	if err != nil {
		return nil
	}
	return &newRule
}

func Filter(s []model.Rule, fn func(rule model.Rule) bool) (match []model.Rule) {
	for _, rule := range s {
		if fn(rule) {
			match = append(match, rule)
		}
	}
	return
}
