package cache

import (
	"github.com/go-redis/redis"
	"grayRelease/src/conf"
)

var redisWhitelistCache RedisWhiteListCache

type RedisWhiteListCache struct {
	Cli *redis.Client
}

func (r RedisWhiteListCache) AddDeviceIdToWhiteList(ruleId int, deviceIds []string) bool {
	e := r.Cli.SAdd(string(ruleId), deviceIds)
	return e.Val() >= 1
}

func (r RedisWhiteListCache) ExpireDeviceIdFromWhiteList(ruleId int, deviceIds []string) bool {
	e := r.Cli.SRem(string(ruleId), deviceIds)
	return e.Val() >= 1
}

func (r RedisWhiteListCache) ExistsDeviceIdInWhiteList(ruleId int, deviceId string) bool {
	e := r.Cli.SIsMember(string(ruleId), deviceId)
	return e.Val()
}

func init() {
	var redisConfig conf.RedisConfig
	redisConfig = conf.GetConfig().RedisWhitelistCache
	redisWhitelistCache.Cli = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
}

func GetRedisWhitelistCache() WhiteListCache {
	return redisWhitelistCache
}
