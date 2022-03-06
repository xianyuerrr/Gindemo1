package cache

import (
	"testing"
)

func Test_getCache(t *testing.T) {
	GetRedisCache()
}

func TestRedis_Hit(t *testing.T) {
	// redisCache := GetRedisCache()
	// client := Client{
	// 	Version:           "1.1",
	// 	DevicePlatform:    "apple",
	// 	DeviceId:          "oppofindx3pro",
	// 	OsApi:             9,
	// 	Channel:           "oppo",
	// 	VersionCode:       "1.1.1.1",
	// 	UpdateVersionCode: "2.2.2.2",
	// 	Aid:               1,
	// 	CpuArch:           "32",
	// }
	// rule := GetRuleById(0)
	// fmt.Println(client)
	// fmt.Println(rule)
}
