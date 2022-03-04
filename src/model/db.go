package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

var Db *gorm.DB

type MysqlConfig struct {
	Dsn      string `yaml:"dsn"`
	MaxOpen  int    `yaml:"max_open"`
	MaxIdle  int    `yaml:"max_idle"`
	LifeTime string `yaml:"life_time"`
	Log      bool   `yaml:"log"`
}

func init() {
	config := MysqlConfig{}

	yamlFile, err := ioutil.ReadFile("config.yml")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &config)

	Db, err = gorm.Open("mysql", config.Dsn)
	if err != nil {
		panic(err)
	}

	sqlDb := Db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(config.MaxIdle)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDb.SetMaxOpenConns(config.MaxOpen)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Hour)
}
