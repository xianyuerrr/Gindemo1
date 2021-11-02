package model

type Rule struct {
	Platform             string `form:"platform" binding:"required"`
	UpdateVersionCode    string `form:"update_version_code" binding:"required"`
	Md5                  string `form:"md_5" binding:"required"`
	DeviceIdList         string `form:"device_id_list" binding:"required"`
	MaxUpdateVersionCode string `form:"max_update_version_code" binding:"required"`
	MinUpdateVersionCode string `form:"min_update_version_code" binding:"required"`
	MaxOsApi             int    `form:"max_os_api" binding:"required"`
	MinOsApi             int    `form:"min_os_api" binding:"required"`
	CpuArch              string `form:"cpu_arch" binding:"required"`
	Channel              string `form:"channel" binding:"required"`
	Title                string `form:"title" binding:"required"`
	UpdateTips           string `form:"update_tips" binding:"required"`
}

func GetRules() *[]Rule {

	//todo connect to Dao
	var rules []Rule

	return &rules
}

func AddRule(r Rule) int {
	//todo

	//return status code
	return 0
}
