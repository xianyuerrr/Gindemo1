package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"grayRelease/src/conf"
	"grayRelease/src/model"
	"math/rand"
	"time"
)

var redisCache RedisCache

type RedisCache struct {
	Cli *redis.Client
}

func (redis RedisCache) Hit(client *model.Client) *string {
	ruleStr, err := redis.Cli.Get(client.String()).Result()
	if err != nil {
		fmt.Println("name does not exist")
		return nil
	}
	return &ruleStr
}

func (redis RedisCache) Store(client *model.Client, res string) bool {
	err := redis.Cli.Set(client.String(), res, time.Hour+(time.Duration(rand.Intn(30))*time.Minute))
	return err == nil
}

func init() {
	initClient()
}

func initClient() (err error) {
	var redisConfig conf.RedisConfig
	redisConfig = conf.GetRedisCacheConfig()

	redisCache.Cli = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	_, err = redisCache.Cli.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedisCache() Cache {
	return redisCache
}
