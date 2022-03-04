package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"grayRelease/src/model"
	"math/rand"
	"time"
)

type Cache interface {
	Hit(client *model.Client) *model.Rule
	Store(client *model.Client, rule *model.Rule) bool
}

type Redis struct {
	Cli *redis.Client
}

func (redis Redis) Hit(client *model.Client) *model.Rule {
	ruleStr, err := redis.Cli.Get(client.String()).Result()
	if err != nil {
		fmt.Println("name does not exist")
		return nil
	}
	return model.CreateRuleFromString(ruleStr)
}

func (redis Redis) Store(client *model.Client, rule *model.Rule) bool {
	err := redis.Cli.Set(client.String(), rule.String(), time.Hour+(time.Duration(rand.Intn(30))*time.Minute))
	return err == nil
}

var cache Redis

func init() {
	initClient()
}

func initClient() (err error) {
	cache.Cli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	res, err := cache.Cli.Ping().Result()
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}

func GetCache() Cache {
	return cache
}
