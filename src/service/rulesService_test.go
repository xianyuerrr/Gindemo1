package service

import (
	"fmt"
	"testing"
)

func TestGetAllRules(t *testing.T) {
	rules := GetAllRules()
	for i := 0; i < len(rules); i++ {
		fmt.Println(rules[i])
	}
}
