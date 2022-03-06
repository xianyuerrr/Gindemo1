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
	mysqlMaster MysqlConfig `yaml:"mysql_master"`
	mysqlFrom   MysqlConfig `yaml:"mysql_from"`
	mysqlAuthor MysqlConfig `yaml:"mysql_author"`
	redisCache  RedisConfig `yaml:"redis_cache"`
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
	return config.mysqlMaster
}

func GetMysqlFromConfig() MysqlConfig {
	return config.mysqlFrom
}

func GetMysqlAuthorConfig() MysqlConfig {
	return config.mysqlAuthor
}

func GetRedisCacheConfig() RedisConfig {
	return config.redisCache
}
