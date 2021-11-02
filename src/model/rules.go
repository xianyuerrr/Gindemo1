package model

type Rule struct {
	Platform             string `form:"platform" binding:"required" json:"platform"`
	UpdateVersionCode    string `form:"update_version_code" binding:"required" json:"update_version_code"`
	Md5                  string `form:"md_5" binding:"required" json:"md_5"`
	DeviceIdList         string `form:"device_id_list" binding:"required" json:"device_id_list"`
	MaxUpdateVersionCode string `form:"max_update_version_code" binding:"required" json:"max_update_version_code"`
	MinUpdateVersionCode string `form:"min_update_version_code" binding:"required" json:"min_update_version_code"`
	MaxOsApi             int    `form:"max_os_api" binding:"required" json:"max_os_api"`
	MinOsApi             int    `form:"min_os_api" binding:"required" json:"min_os_api"`
	CpuArch              string `form:"cpu_arch" binding:"required" json:"cpu_arch"`
	Channel              string `form:"channel" binding:"required" json:"channel"`
	Title                string `form:"title" binding:"required" json:"title"`
	UpdateTips           string `form:"update_tips" binding:"required" json:"update_tips"`
}

func GetRules() *[]Rule {

	//todo connect to Dao
	var rules []Rule

	return &rules
}

func AddRule(r Rule) bool {
	//todo

	//return status
	return true
}
