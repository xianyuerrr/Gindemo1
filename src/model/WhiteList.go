package model

// WhiteList 设备ID白名单
type WhiteList struct {
	RuleID   int    `gorm:"rule_id" json:"rule_id" label:"这个白名单对应的规则的id"`
	DeviceID string `gorm:"device_id" json:"device_id" label:"设备ID"`
}

func (d *WhiteList) TableName() string {
	return "whiteLists"
}
