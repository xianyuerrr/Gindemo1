package service

import "grayRelease/src/model"

// GetAllRules 获取所有未删除的rule
func GetAllRules() []model.Rule {
	return nil
}

// AddRule 新增rule，若已存在aid相同的，将其启用且更新
func AddRule(rule *model.Rule) bool {
	return true
}

// UpdateRule 更新rule
func UpdateRule(rule *model.Rule) bool {
	return true
}

// DeleteRule 删除rule（不是真的从数据库删除）
func DeleteRule(aid int) bool {
	return true
}

// ReleaseRule 上线版本
func ReleaseRule() bool {
	return true
}

// OfflineRule 下线版本
func OfflineRule() bool {
	return true
}



// GetReleasedRules 获取以及发布上线的rule
func GetReleasedRules() []model.Rule {
	return nil
}
