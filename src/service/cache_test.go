package service

import (
	"fmt"
	"grayRelease/src/model"
	"testing"
)

func Test_getCache(t *testing.T) {
	GetCache()
}

func TestRedis_Hit(t *testing.T) {
	cache := GetCache()
	client := model.Client{
		Version:           "1.1",
		DevicePlatform:    "apple",
		DeviceId:          "oppofindx3pro",
		OsApi:             9,
		Channel:           "oppo",
		VersionCode:       "1.1.1.1",
		UpdateVersionCode: "2.2.2.2",
		Aid:               1,
		CpuArch:           "32",
	}
	rule := model.GetRuleById(0)
	fmt.Println(client)
	fmt.Println(rule)

	cache.Store(&client, rule.Rule)
	r := cache.Hit(&client)
	fmt.Println(r)
	fmt.Println(r == rule.Rule)
}
