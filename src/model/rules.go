package model

import (
	"github.com/jmoiron/sqlx"
)

type Rule struct {
	Aid                  int    `json:"aid" form:"aid" binding:"required" db:"aid"`
	Platform             string `form:"platform" binding:"required" json:"platform" db:"platform"`
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

var Db *sqlx.DB

//func init() {
//	database, err := sqlx.Open("mysql", "root:kgkg@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//	Db = database
//	defer Db.Close() // 注意这行代码要写在上面err判断的下面
//}

func CheckRule(client Client) Rule {

	//todo connect to Database
	var ruleT Rule

	return ruleT
}

func AddRule(rule Rule) bool {
	//todo
	//r, err := Db.Exec("insert into rule(platform, update_version_code, md_5, device_id_list) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	//	"stu001", "man", "stu01@qq.com")
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return false
	//}
	//id, err := r.LastInsertId()
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return false
	//}
	//
	//fmt.Println("insert succ:", id)
	//return status
	return true
}
