package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"grayRelease/src/conf"
	"grayRelease/src/model"
	"math/rand"
	"time"
)

var redisCheckCache RedisCheckCache

type RedisCheckCache struct {
	Cli *redis.Client
}

func (redis RedisCheckCache) Hit(client *model.Client) *string {
	ruleStr, err := redis.Cli.Get(client.String()).Result()
	if err != nil {
		fmt.Println("name does not exist")
		return nil
	}
	return &ruleStr
}

func (redis RedisCheckCache) Store(client *model.Client, res string) bool {
	// fixme: 将其反转的话，方便在 offline rule 时将其无效化。但是这样就不方便为 check 请求设置过期时间了
	err := redis.Cli.Set(client.String(), res, time.Hour+(time.Duration(rand.Intn(30))*time.Minute))
	return err == nil
}

func init() {
	var redisConfig conf.RedisConfig
	redisConfig = conf.GetRedisCacheConfig()
	redisCheckCache.Cli = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
}

func GetCheckRedisCache() CheckCache {
	return redisCheckCache
}
