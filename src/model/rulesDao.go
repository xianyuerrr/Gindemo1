package model

import (
	"time"
)

func NewRule2Rule(rule NewRule) Rule {
	return *rule.Rule
}

func NewRules2Rules(rules []NewRule) []Rule {
	n := len(rules)
	res := make([]Rule, n)
	for i := 0; i < n; i++ {
		res[i] = NewRule2Rule(rules[i])
	}
	return res
}

func rule2NewRule(rule *Rule) NewRule {
	return NewRule{Rule: rule,
		CreatTime:  time.Now(),
		DeleteTime: time.Now(),
		IsDelete:   0,
		IsRelease:  0}
}

func Rules2NewRules(rules []Rule) []NewRule {
	n := len(rules)
	res := make([]NewRule, n)
	for i := 0; i < n; i++ {
		res[i] = rule2NewRule(&rules[i])
	}
	return res
}

// GetAllRules 获取数据库中所有 is_delete 为 0 的记录
func GetAllRules() []NewRule {
	var rules []NewRule
	db := GetReadDb()
	db.Where("is_delete = ?", 0).Find(&rules)
	return rules
}

func GetReleasedRules(aid int, platform string, channel string) []NewRule {
	var rules []NewRule
	db := GetReadDb()
	db.Debug().
		Where("aid = ? AND is_release = ? AND platform = ? AND channel = ?", aid, 1, platform, channel).
		Find(&rules)
	return rules
}

func GetRuleById(id uint) *NewRule {
	var newRule NewRule
	db := GetReadDb()
	err := db.First(&newRule, id)
	if err.Error != nil {
		return nil
	}
	return &newRule
}

func AddRule(rule *Rule) bool {
	var newRule = rule2NewRule(rule)
	db := GetWriteDb()
	db.AutoMigrate(&NewRule{})
	err := db.Create(&newRule)
	return err.Error == nil
}

// RemoveRule 删除
func RemoveRule(id uint) bool {
	db := GetWriteDb()
	err := db.Delete(&NewRule{}, id)
	return err.Error == nil
}

func UpdateRule(newRule *NewRule) bool {
	// err := masterDb.Model(&NewRule{}).Select("*").Omit("id").Update(newRule)
	db := GetReadDb()
	err := db.Save(newRule)
	return err.RowsAffected == 1
}
