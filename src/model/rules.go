package model

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// todo test

type Rule struct {
	// app的唯⼀标识，全匹配
	Aid int `form:"aid" binding:"required" json:"aid" gorm:"Column:aid"`
	// 平台，Android/iOS，字符串匹配
	Platform string `form:"platform" binding:"required" json:"platform" db:"platform" gorm:"Column:platform; Type:varchar(20)"`
	// 命中后的包下载链接，命中后返回
	DownloadUrl string `form:"download_url" binding:"required" json:"download_url" db:"download_url" gorm:"Column:download_url; Type:varchar(128)" `
	// 包的当前版本号例如（8.1.3.01），命中后返回
	UpdateVersionCode string `form:"update_version_code" binding:"required" json:"update_version_code" gorm:"Column:update_version_code; Type:varchar(128)"`
	// 包的MD5，命中后返回
	Md5 string `form:"md_5" binding:"required" json:"md_5" gorm:"Column:md5; Type:varchar(128)"`
	// 设备号⽩名单列表在⽩名单范围内即为命中
	DeviceIdList string `form:"device_id_list" binding:"required" json:"device_id_list" gorm:"Column:device_id_list; Type:text"`
	// 可升级的最⼤版本号上传的版本号需⼩于或等于配置版本号
	MaxUpdateVersionCode string `form:"max_update_version_code" binding:"required" json:"max_update_version_code" gorm:"Column:max_update_version_code; Type:varchar(128)"`
	// 可升级的最⼩版本号上传的版本号需⼤于或等于配置版本号；
	MinUpdateVersionCode string `form:"min_update_version_code" binding:"required" json:"min_update_version_code" gorm:"Column:min_update_version_code; Type:varchar(128)"`
	// ⽀持的最⼤操作系统版本（适⽤于安卓）上传的操作系统版本号需⼩于或等于配置版本号；
	MaxOsApi int `form:"max_os_api" binding:"required" json:"max_os_api" gorm:"Column:max_os_api"`
	// ⽀持的最⼩操作系统版本（适⽤于安卓）上传的操作系统版本号需⼤于或等于配置版本号；
	MinOsApi int `form:"min_os_api" binding:"required" json:"min_os_api" gorm:"Column:min_os_api"`
	// CPU架构，32/64，32位命中32位版本，64位命中64位版本
	CpuArch string `form:"cpu_arch" binding:"required" json:"cpu_arch" gorm:"Column:cpu_arch; Type:varchar(32)"`
	// 渠道号，字符串匹配
	Channel string `form:"channel" binding:"required" json:"channel" gorm:"Column:channel; Type:varchar(128)"`
	// 弹窗标题，命中后返回
	Title string `form:"title" binding:"required" json:"title" gorm:"Column:title; Type:varchar(256)"`
	// 弹窗的更新⽂本，命中后返回
	UpdateTips string `form:"update_tips" binding:"required" json:"update_tips" gorm:"Column:update_tips; Type:varchar(1024)"`
}

// 新版本检查接⼝规则，多条件的⽐较顺序是：
// 业务id > platform > 渠道 > 设备⽩名单 > 【其他条件计算顺序均可】
// 如果命中，返回满⾜条件的升级包的基本信息；⾄多只能返回⼀条升级包规则；

func (r Rule) String() string {
	buf, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(buf)
}

type NewRule struct {
	ID uint `gorm:"Column:id; primary_key; AUTO_INCREMENT" json:"id"`
	*Rule
	CreatTime  time.Time `gorm:"Column:create_time" json:"creat_time"`
	DeleteTime time.Time `gorm:"Column:delete_time" json:"delete_time"`
	IsDelete   int       `gorm:"Column:is_delete" json:"is_delete"`
	IsRelease  int       `gorm:"Column:is_release" json:"is_release"`
}

func (r NewRule) TableName() string {
	return "rules"
}

func (r NewRule) String() string {
	buf, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(buf)
}

func rule2NewRule(rule Rule) NewRule {
	return NewRule{Rule: &rule,
		CreatTime:  time.Now(),
		DeleteTime: time.Now(),
		IsDelete:   0,
		IsRelease:  0}
}

// GetAllRules 根据 deviceId 获取对应的 rules
func GetAllRules() *[]Rule {
	var newRules []NewRule
	Db.Find(&newRules)
	rules := make([]Rule, 0)
	for index := 0; index < len(newRules); index++ {
		rules = append(rules, *(newRules[index].Rule))
	}
	return &rules
}

func GetRuleById(id uint) *NewRule {
	var newRule NewRule
	err := Db.First(&newRule, id)
	if err.Error != nil {
		return nil
	}
	return &newRule
}

func GetRuleByAid(aid int) *NewRule {
	var newRule NewRule
	err := Db.Where("aid = ?", aid).First(&newRule)
	if err.Error != nil {
		fmt.Println(err.Error)
		return nil
	}
	return &newRule
}

func AddRule(rule *Rule) bool {
	var newRule = rule2NewRule(*rule)
	Db.AutoMigrate(&NewRule{})
	err := Db.Create(&newRule)
	return err.Error == nil
}

// RemoveRule 删除
func RemoveRule(id uint) bool {
	err := Db.Delete(&NewRule{}, id)
	if err.Error != nil {
		fmt.Println(err.Error)
		return false
	}
	return true
}

func UpdateRule(newRule *NewRule) bool {
	err := Db.Model(&NewRule{}).Where("id = ?", newRule.ID).Update(newRule)
	return err.RowsAffected == 1
}
