package model

import (
	"fmt"
	"testing"
)

func TestGetRules(t *testing.T) {
	rules := GetAllRules()
	for i := 0; i < len(*rules); i++ {
		fmt.Println((*rules)[i])
	}
}

func Test_rule2newRule(t *testing.T) {

}

func TestAddRule(t *testing.T) {
	rule := Rule{
		Aid:                  0,
		Platform:             "platform",
		DownloadUrl:          "download_url",
		UpdateVersionCode:    "update_version_code",
		Md5:                  "md5",
		DeviceIdList:         "device_id_list",
		MaxUpdateVersionCode: "max_update_version_code",
		MinUpdateVersionCode: "min_update_version_code",
		MaxOsApi:             0,
		MinOsApi:             0,
		CpuArch:              "x64",
		Channel:              "huawei",
		Title:                "title",
		UpdateTips:           "update_tips",
	}
	i := AddRule(&rule)
	fmt.Println(i)
}

func TestGetRulesById(t *testing.T) {
	rule := GetRuleById(1)
	fmt.Println(rule)
}

func TestGetRuleByAid(t *testing.T) {
	r := GetRuleByAid(2)
	fmt.Println(r)
}

func TestRemoveRule(t *testing.T) {
	fmt.Println(RemoveRule(1))
}

func TestUpdateRule(t *testing.T) {
	rule := GetRuleById(0)
	fmt.Println(rule)
	rule.Aid = 4
	rule.IsDelete = 1
	// rule.ID = 20
	fmt.Println(rule)
	fmt.Println(UpdateRule(rule))
}
