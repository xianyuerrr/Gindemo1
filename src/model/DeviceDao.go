package model

// AddWhiteList 添加向规则白名单添加设备
func AddWhiteList(ruleId int, deviceId string) bool {
	var whiteList WhiteList
	whiteList = WhiteList{
		RuleID:   ruleId,
		DeviceID: deviceId,
	}
	db := GetWriteDb()
	db.AutoMigrate(&whiteList)
	err := db.Create(&whiteList)
	return err.Error == nil
}

// DeleteWhiteList 从规则白名单删除设备
func DeleteWhiteList(ruleId int, deviceIds []string) bool {
	var whiteLists WhiteList
	db := GetWriteDb()
	err := db.Where("rule_id = ? AND device_id IN ?", ruleId, deviceIds).
		Delete(&whiteLists)
	return err.Error == nil
}

// ExistsInWhiteList 查找设备是否在规则的白名单中
func ExistsInWhiteList(ruleId int, deviceId string) bool {
	var whiteLists []WhiteList
	db := GetReadDb()
	db.Where("rule_id = ? And device_id = ?", ruleId, deviceId).
		Find(&whiteLists)
	return len(whiteLists) > 0
}

// GetDeviceIdsByRuleId 根据 ruleId 获取对应的白名单
func GetDeviceIdsByRuleId(ruleId int) []string {
	var whiteLists []WhiteList
	db := GetReadDb()
	db.Where("rule_id = ?", ruleId).
		Find(&whiteLists)
	n := len(whiteLists)
	deviceIds := make([]string, n)
	for i := 0; i < n; i++ {
		deviceIds[i] = whiteLists[i].DeviceID
	}
	return deviceIds
}

// GetRuleIdsByDeviceId 根据 aid 和 deviceId 获取白名单包含其的规则 id
func GetRuleIdsByDeviceId(deviceId string) []int {
	var whiteLists []WhiteList
	db := GetReadDb()
	db.Where("device_id = ?", deviceId).
		Find(&whiteLists)
	n := len(whiteLists)
	ruleIds := make([]int, n)
	for i := 0; i < n; i++ {
		ruleIds[i] = whiteLists[i].RuleID
	}
	return ruleIds
}

// ruleId 唯一，其对应规则的白名单中包含多个 deviceId
// check 的时候，我们使用 deviceId 和 aid 获取 ruleId

// 由于一条 rule 的白名单包含 deviceId 的数目非常多，可达 10w+
// 而且 rule 具有时效性，所以以 deviceId 和 aid 进行查询所得的 rule 数目会比较有限

// // CacheDeviceID 将设备白名单放到缓存里面
// func CacheDeviceID(data *model.Device) {
// 	key := "app_device_id_" + strconv.Itoa(data.RuleID)
// 	redis.RedisClient.SAdd(context.Background(), key, data.DeviceID)
// 	redis.RedisClient.Expire(context.Background(), key, config.RedisSetting.ExpireTime)
// }
