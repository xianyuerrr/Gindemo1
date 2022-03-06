package model

import (
	"fmt"
	"testing"
)

func TestGetRules(t *testing.T) {
	rules := GetAllRules()
	for i := 0; i < len(rules); i++ {
		fmt.Println((rules)[i])
	}
}

func Test_rule2newRule(t *testing.T) {

}

func TestAddRule(t *testing.T) {
	// current := time.Now()
	// fmt.Println(current.Unix())
	// fmt.Println(current.Add(time.Minute).Unix())
	// fmt.Println(time.Minute.Milliseconds())
	AddRule(&Rule{})
}

func TestGetRulesById(t *testing.T) {
	rule := GetRuleById(1)
	fmt.Println(rule)
}

func TestGetRuleByAid(t *testing.T) {

}

func TestRemoveRule(t *testing.T) {
	fmt.Println(RemoveRule(1))
}

func TestUpdateRule(t *testing.T) {
	rule := GetRuleById(1)
	fmt.Println(rule)
	rule.Aid = 4
	rule.IsDelete = 1
	// rule.Id = 3
	fmt.Println(rule)
	fmt.Println(UpdateRule(rule))
}
