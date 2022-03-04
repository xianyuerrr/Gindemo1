package model

import (
	"fmt"
	"time"
)

func rule2NewRule(rule Rule) NewRule {
	return NewRule{Rule: &rule,
		CreatTime:  time.Now(),
		DeleteTime: time.Now(),
		IsDelete:   0,
		IsRelease:  0}
}

// GetAllRules 根据 deviceId 获取对应的 rules
func GetAllRules() *[]NewRule {
	var rules []NewRule
	Db.Find(&rules)
	return &rules
}

func GetRuleById(id uint) *NewRule {
	var newRule NewRule
	err := Db.First(&newRule, id)
	if err.Error != nil {
		return nil
	}
	return &newRule
}

func GetRuleByAid(aid int) *NewRule {
	var newRule NewRule
	err := Db.Where("aid = ?", aid).First(&newRule)
	if err.Error != nil {
		fmt.Println(err.Error)
		return nil
	}
	return &newRule
}

func AddRule(rule *Rule) bool {
	var newRule = rule2NewRule(*rule)
	Db.AutoMigrate(&NewRule{})
	err := Db.Create(&newRule)
	return err.Error == nil
}

// RemoveRule 删除
func RemoveRule(id uint) bool {
	err := Db.Delete(&NewRule{}, id)
	if err.Error != nil {
		fmt.Println(err.Error)
		return false
	}
	return true
}

func UpdateRule(newRule *NewRule) bool {
	err := Db.Model(&NewRule{}).Where("id = ?", newRule.ID).Update(newRule)
	return err.RowsAffected == 1
}
