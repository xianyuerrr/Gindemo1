package model

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Rule struct {
	Aid                  int    `json:"aid" form:"aid" binding:"required" gorm:"Column:aid"`
	Platform             string `form:"platform" binding:"required" json:"platform" db:"platform" gorm:"Column:platform; Type:varchar(128)"`
	DownloadUrl          string `form:"download_url" binding:"required" json:"download_url" db:"download_url" gorm:"Column:download_url; Type:varchar(1024)" `
	UpdateVersionCode    string `form:"update_version_code" binding:"required" json:"update_version_code" gorm:"Column:update_version_code; Type:varchar(128)"`
	Md5                  string `form:"md_5" binding:"required" json:"md_5" gorm:"Column:md5; Type:varchar(128)"`
	DeviceIdList         string `form:"device_id_list" binding:"required" json:"device_id_list" gorm:"Column:device_id_list; Type:text"`
	MaxUpdateVersionCode string `form:"max_update_version_code" binding:"required" json:"max_update_version_code" gorm:"Column:max_update_version_code; Type:varchar(128)"`
	MinUpdateVersionCode string `form:"min_update_version_code" binding:"required" json:"min_update_version_code" gorm:"Column:min_update_version_code; Type:varchar(128)"`
	MaxOsApi             int    `form:"max_os_api" binding:"required" json:"max_os_api" gorm:"Column:max_os_api"`
	MinOsApi             int    `form:"min_os_api" binding:"required" json:"min_os_api" gorm:"Column:min_os_api"`
	CpuArch              string `form:"cpu_arch" binding:"required" json:"cpu_arch" gorm:"Column:cpu_arch; Type:varchar(32)"`
	Channel              string `form:"channel" binding:"required" json:"channel" gorm:"Column:channel; Type:varchar(128)"`
	Title                string `form:"title" binding:"required" json:"title" gorm:"Column:title; Type:varchar(256)"`
	UpdateTips           string `form:"update_tips" binding:"required" json:"update_tips" gorm:"Column:update_tips; Type:varchar(1024)"`
}

var Db *gorm.DB

func init() {
	dsn := "root:wang0805@tcp(127.0.0.1:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db = db
}

//func GetRules() *[]Rule {
//
//}

func AddRule(rule Rule) bool {
	Db.AutoMigrate(&Rule{})
	if err := Db.Create(&rule).Error; err != nil {
		return false
	}
	return true
}
