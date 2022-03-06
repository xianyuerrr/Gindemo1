package model

import (
	"fmt"
	"grayRelease/src/model/tables"
	"time"
)

func NewRule2Rule(rule tables.NewRule) tables.Rule {
	return *rule.Rule
}

func NewRules2Rules(rules []tables.NewRule) []tables.Rule {
	n := len(rules)
	res := make([]tables.Rule, n)
	for i := 0; i < n; i++ {
		res[i] = NewRule2Rule(rules[i])
	}
	return res
}

func rule2NewRule(rule tables.Rule) tables.NewRule {
	return tables.NewRule{Rule: &rule,
		CreatTime:  time.Now(),
		DeleteTime: time.Now(),
		IsDelete:   0,
		IsRelease:  0}
}

func Rules2NewRules(rules []tables.Rule) []tables.NewRule {
	n := len(rules)
	res := make([]tables.NewRule, n)
	for i := 0; i < n; i++ {
		res[i] = rule2NewRule(rules[i])
	}
	return res
}

// GetAllRules 获取数据库中所有 is_delete 为 0 的记录
func GetAllRules() []tables.NewRule {
	var rules []tables.NewRule
	db := GetReadDb()
	db.Where("is_delete = ?", 0).Find(&rules)
	return rules
}

func GetReleasedRules(aid int) []tables.NewRule {
	var rules []tables.NewRule
	db := GetReadDb()
	db.Debug().Where("aid = ? AND is_release = ?", aid, 1).Find(&rules)
	return rules
}

func GetRuleById(id uint) *tables.NewRule {
	var newRule tables.NewRule
	db := GetReadDb()
	err := db.First(&newRule, id)
	if err.Error != nil {
		return nil
	}
	return &newRule
}

func AddRule(rule *tables.Rule) bool {
	var newRule = rule2NewRule(*rule)
	db := GetWriteDb()
	db.AutoMigrate(&tables.NewRule{})
	err := masterDb.Create(&newRule)
	return err.Error == nil
}

// RemoveRule 删除
func RemoveRule(id uint) bool {
	db := GetWriteDb()
	err := db.Delete(&tables.NewRule{}, id)
	if err.Error != nil {
		fmt.Println(err.Error)
		return false
	}
	return true
}

func UpdateRule(newRule *tables.NewRule) bool {
	// err := masterDb.Model(&NewRule{}).Select("*").Omit("id").Update(newRule)
	db := GetReadDb()
	err := db.Save(newRule)
	return err.RowsAffected == 1
}
