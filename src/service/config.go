package service

import (
	"demo1/src/model"
)

func CheckConfig(config model.Rule) bool {
	//todo
	//todo 版本号化简
	return model.AddRule(config)
}
