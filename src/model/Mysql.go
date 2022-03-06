package model

import (
	"github.com/jinzhu/gorm"
	"grayRelease/src/conf"
	"time"
)

// 准备项目使用到的数据库

var masterDb *gorm.DB
var fromDb *gorm.DB

func getDb(dsn string, maxOpen int, maxIdle int) *gorm.DB {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
		return nil
	}
	sqlDb := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(maxIdle)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDb.SetMaxOpenConns(maxOpen)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db
}

// GetWriteDb 返回 配置端 用到的数据库
func GetWriteDb() *gorm.DB {
	if masterDb == nil {
		masterConfig := conf.GetMysqlMasterConfig()
		masterDb = getDb(masterConfig.Dsn, masterConfig.MaxOpen, masterConfig.MaxIdle)
	}
	return masterDb
}

// GetReadDb 返回 check 端用到的数据库
func GetReadDb() *gorm.DB {
	if fromDb == nil {
		fromConfig := conf.GetMysqlFromConfig()
		fromDb = getDb(fromConfig.Dsn, fromConfig.MaxOpen, fromConfig.MaxIdle)
	}
	return fromDb.Debug()
}
