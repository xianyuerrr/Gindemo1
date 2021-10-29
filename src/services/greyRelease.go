package services

import (
	"demo1/src/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func Hit(c *gin.Context) {
	var respUrl string
	appVerion := c.Query("appVersion")
	userId := c.Query("userDID")
	rules := model.GetRules()

	for index := 0; index < len(*rules); index++ {
		if cast.ToInt(userId) < (*rules)[index].MinUserDID || cast.ToInt(userId) > (*rules)[index].MaxUserDID {
			continue
		}

		if cast.ToInt(appVerion) < (*rules)[index].MinVersion || cast.ToInt(appVerion) > (*rules)[index].MaxVersion {
			continue
		}

		respUrl = (*rules)[index].GreyLink
	}
	c.JSON(http.StatusOK, gin.H{"url:": respUrl})
}
