package model

import "github.com/jinzhu/gorm"

type WhiteList struct {
	gorm.Model
	DeviceName string    `gorm:"Column:device_name" json:"device_name"`
	rules      []NewRule `gorm:"many2many:rule2device_list;"`
}

func (WhiteList) TableName() string {
	return "white_list"
}
func ParseWhiteList(DeviceIdList string) []WhiteList {
	//todo
	return []WhiteList{}
}
