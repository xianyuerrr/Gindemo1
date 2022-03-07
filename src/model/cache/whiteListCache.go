package cache

type WhiteListCache interface {
	AddDeviceIdToWhiteList(ruleId int, deviceId []string) bool
	ExpireDeviceIdFromWhiteList(ruleId int, deviceId []string) bool
	ExistsDeviceIdInWhiteList(ruleId int, deviceId string) bool
}
