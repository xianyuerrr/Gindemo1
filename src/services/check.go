package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type client struct {
	Version           string `json:"version" binding:"required"`
	DevicePlatform    string `json:"device_platform" binding:"required"`
	DeviceId          string `json:" device_id" binding:"required"`
	OsApi             string `json:"os_api" binding:"required"`
	Channel           string `json:"channel" binding:"required"`
	VersionCode       string `json:"version_code" binding:"required"`
	UpdateVersionCode string `json:"update_version_code" binding:"required"`
	Aid               string `json:"aid" binding:"required"`
}

func hit(form client) (string, string, string, string, string) {
	//todo

	//if hit,
	//return downloadUrl, updateVersionCode, md5, title, updateTips

	return "11", "22", "33", "44", "55"
}
func Check(c *gin.Context) {
	var form client
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	downloadUrl, updateVersionCode, md5, title, updateTips := hit(form)
	c.JSON(http.StatusOK, gin.H{
		"download_url:":       downloadUrl,
		"update_version_code": updateVersionCode,
		"md5":                 md5,
		"title":               title,
		"update_tips":         updateTips,
	})
}
