package services

import (
	"demo1/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Check(c *gin.Context) {
	var respUrl string
	version := c.Query("version")
	device_platform := c.Query("device_platform")
	device_id := c.Query("device_id")
	os_api := c.Query("os_api")
	channel := c.Query("channel")
	version_code := c.Query("version_code")
	update_version_code := c.Query("update_version_code")
	aid := c.Query("aid")
	cpu_arch := c.Query("cpu_arch")
	rules := model.GetRules()

	for index := 0; index < len(*rules); index++ {
		//if rule continue
		respUrl = (*rules)[index].GreyLink
		break
	}
	c.JSON(http.StatusOK, gin.H{"url:": respUrl})
}
