package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type MysqlConfig struct {
	Dsn      string `yaml:"dsn"`
	MaxOpen  int    `yaml:"max_open"`
	MaxIdle  int    `yaml:"max_idle"`
	LifeTime string `yaml:"life_time"`
	Log      bool   `yaml:"log"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Config struct {
	MysqlMaster         MysqlConfig `yaml:"mysql_master"`
	MysqlFrom           MysqlConfig `yaml:"mysql_from"`
	MysqlAuthor         MysqlConfig `yaml:"mysql_author"`
	RedisCheckCache     RedisConfig `yaml:"redis_check_cache"`
	RedisWhitelistCache RedisConfig `yaml:"redis_whitelist_cache"`
}

var config Config

func init() {
	yamlFile, err := ioutil.ReadFile("./src/conf/config.yml")

	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("加载配置失败")
		panic(err)
	}
}

func GetConfig() Config {
	return config
}

func GetMysqlMasterConfig() MysqlConfig {
	return config.MysqlMaster
}

func GetMysqlFromConfig() MysqlConfig {
	return config.MysqlFrom
}

func GetMysqlAuthorConfig() MysqlConfig {
	return config.MysqlAuthor
}

func GetRedisCacheConfig() RedisConfig {
	return config.RedisCheckCache
}
