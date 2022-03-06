package authencator

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// 简单单例
// var mysqlCredentialStorage *MysqlCredentialStorage

// MysqlCredentialStorage CredentialStorage 的默认实现类
type MysqlCredentialStorage struct {
	authorDb *gorm.DB
	dsn      string
	maxOpen  int
	maxIdle  int
}

// AddAuthor 添加用户
func (mysqlCredentialStorage MysqlCredentialStorage) AddAuthor(author *Author) bool {
	mysqlCredentialStorage.authorDb.AutoMigrate(&Author{})
	err := mysqlCredentialStorage.authorDb.Create(author)
	return err.Error == nil
}

// GetPasswordByUserId 根据 userId 查询 password
func (mysqlCredentialStorage MysqlCredentialStorage) GetPasswordByUserId(userId string) string {
	var author Author
	err := mysqlCredentialStorage.authorDb.Where("user_id = ?", userId).First(&author)
	if err.Error != nil {
		return ""
	}
	return author.Password
}

// UpdateAuthor 更新密码
func (mysqlCredentialStorage MysqlCredentialStorage) UpdateAuthor(author Author) bool {
	err := mysqlCredentialStorage.authorDb.Save(author)
	return err == nil
}

// DeleteAuthor 删除用户
func (mysqlCredentialStorage MysqlCredentialStorage) DeleteAuthor(userId string) bool {
	err := mysqlCredentialStorage.authorDb.Delete(&Author{}, userId)
	return err.Error == nil
}

// 为单例 mysqlCredentialStorage 创建对象赋值
func getDb(dsn string, maxIdle int, maxOpen int) *gorm.DB {
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
	sqlDb.SetConnMaxLifetime(time.Minute)
	return db.Debug()
}

// GetMysqlCredentialStorage 传入参数获取 MysqlCredentialStorage 对象
func GetMysqlCredentialStorage(dsn string, maxIdle int, maxOpen int) *MysqlCredentialStorage {
	return &MysqlCredentialStorage{authorDb: getDb(dsn, maxIdle, maxOpen)}
}
