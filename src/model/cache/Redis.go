package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"grayRelease/src/conf"
	"grayRelease/src/model/tables"
	"math/rand"
	"time"
)

var cache RedisCache

type RedisCache struct {
	Cli *redis.Client
}

func (redis RedisCache) Hit(client *tables.Client) *string {
	ruleStr, err := redis.Cli.Get(client.String()).Result()
	if err != nil {
		fmt.Println("name does not exist")
		return nil
	}
	return &ruleStr
}

func (redis RedisCache) Store(client *tables.Client, res string) bool {
	err := redis.Cli.Set(client.String(), res, time.Hour+(time.Duration(rand.Intn(30))*time.Minute))
	return err == nil
}

func init() {
	initClient()
}

func initClient() (err error) {
	var redisConfig conf.RedisConfig
	redisConfig = conf.GetRedisCacheConfig()

	cache.Cli = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	_, err = cache.Cli.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetCache() Cache {
	return cache
}
