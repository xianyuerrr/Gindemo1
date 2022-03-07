package service

import (
	"grayRelease/src/model"
)

// GetAllRules 获取所有 is_delete 为 0 的rule
func GetAllRules() []model.Rule {
	rules := model.GetAllRules()
	return model.NewRules2Rules(rules)
}

// AddRule 新增rule，若已存在aid相同的，将其启用且更新
func AddRule(rule *model.Rule) bool {
	newRule := model.GetRuleById(rule.Id)
	// 不存在此记录，将其添加进数据库
	if newRule == nil {
		// 防止自增主键跳跃插入导致主键值过大
		rule.Id = 0
		return model.AddRule(rule)
	}
	// 存在此纪录，更新其 is_delete 字段以及 Rule
	newRule.IsDelete = 0
	model.UpdateRule(newRule)
	return true
}

// UpdateRule 更新rule
func UpdateRule(rule *model.Rule) bool {
	newRule := model.GetRuleById(rule.Id)
	// 不存在此记录，无法更新
	if newRule == nil {
		return false
	}
	// 存在此纪录，更新 Rule
	newRule.Rule = rule
	model.UpdateRule(newRule)
	return true
}

// DeleteRule 删除rule（不是真的从数据库删除）
func DeleteRule(id uint) bool {
	newRule := model.GetRuleById(id)
	if newRule == nil {
		return false
	}
	newRule.IsDelete = 1
	newRule.IsRelease = 0
	return model.UpdateRule(newRule)
}

// ReleaseRule 上线版本
func ReleaseRule(id uint) bool {
	newRule := model.GetRuleById(id)
	if newRule == nil {
		return false
	}
	newRule.IsRelease = 1
	return model.UpdateRule(newRule)
}

// OfflineRule 下线版本
func OfflineRule(id uint) bool {
	newRule := model.GetRuleById(id)
	if newRule == nil {
		return false
	}
	newRule.IsRelease = 0
	return model.UpdateRule(newRule)
}

// GetReleasedRules 获取以及发布上线的rule
func GetReleasedRules(aid int, platform string, channel string) []model.Rule {
	rules := model.GetReleasedRules(aid, platform, channel)
	return model.NewRules2Rules(rules)
}
