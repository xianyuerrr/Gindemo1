package service

import (
	"grayRelease/src/model"
	"grayRelease/src/model/cache"
)

// AddDeviceIdToWhiteList AddWhiteList Add 向白名单新增规则
func AddDeviceIdToWhiteList(ruleId int, deviceIds []string) int {
	cnt := 0
	for _, deviceId := range deviceIds {
		// 若白名单中已存在，则跳过
		if ExistsDeviceIdInWhiteList(ruleId, deviceId) {
			continue
		}
		// 若白名单中还没有，则添加
		if model.AddWhiteList(ruleId, deviceId) {
			cnt++
		}
	}
	return cnt
}

// DeleteDeviceIdFromWhiteList 从白名单删除设备
func DeleteDeviceIdFromWhiteList(ruleId int, deviceIds []string) bool {
	// 先从redis里面删除这条ID
	cache := cache.GetRedisWhitelistCache()
	cache.ExpireDeviceIdFromWhiteList(ruleId, deviceIds)
	// 然后从数据库里面删除
	return model.DeleteWhiteList(ruleId, deviceIds)
}

// ExistsDeviceIdInWhiteList 查找某一个device_id是否在某一条规则的白名单
func ExistsDeviceIdInWhiteList(ruleId int, deviceId string) bool {
	// 先查询缓存，看是否设备是否在该规则的白名单，若缓存未命中再查询数据库
	cache := cache.GetRedisWhitelistCache()
	if cache.ExistsDeviceIdInWhiteList(ruleId, deviceId) {
		return true
	}
	return model.ExistsInWhiteList(ruleId, deviceId)
}
