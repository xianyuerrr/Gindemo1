package model

import (
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type Config struct {
	Mysql struct {
		Dsn      string `yaml:"dsn"`
		MaxOpen  int    `yaml:"max_open"`
		MaxIdle  int    `yaml:"max_idle"`
		LifeTime string `yaml:"life_time"`
		Log      bool   `yaml:"log"`
	}
}

func Mysql() {
	//dbConfig := Config{}
	//
	//data, err := ioutil.ReadFile("https://github.com/kguniverse/techtrainingcamp-AppUpgrade/blob/master/src/config.yml")
	//if err != nil {
	//	panic(err)
	//}
	//err = yaml.Unmarshal([]byte(data), &dbConfig)
	//if err != nil {
	//	panic(err)
	//}
	//db, err := gorm.Open("mysql", dbConfig.Mysql.Dsn)
	//if err != nil {
	//	panic(err)
	//}
	Dsn := "kgkg:Wang0805@tcp(rm-bp17dut928o4el9fcvo.mysql.rds.aliyuncs.com:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", Dsn)
	if err != nil {
		panic(err)
	}
	Db = db
}
