package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type client struct {
	version           string
	devicePlatform    string
	deviceId          string
	osApi             string
	channel           string
	versionCode       string
	updateVersionCode string
	aid               string
}

func hit(cq map[string][]string) (string, string, string, string, string) {
	//todo

	//if hit,
	//return downloadUrl, updateVersionCode, md5, title, updateTips
	return "11", "22", "33", "44", "55"
}
func Check(c *gin.Context) {
	cq := c.Request.URL.Query()

	//version := c.Query("version")
	//device_platform := c.Query("device_platform")
	//device_id := c.Query("device_id")
	//os_api := c.Query("os_api")
	//channel := c.Query("channel")
	//version_code := c.Query("version_code")
	//update_version_code := c.Query("update_version_code")
	//aid := c.Query("aid")
	//cpu_arch := c.Query("cpu_arch")
	downloadUrl, updateVersionCode, md5, title, updateTips := hit(cq)
	c.JSON(http.StatusOK, gin.H{
		"download_url:":       downloadUrl,
		"update_version_code": updateVersionCode,
		"md5":                 md5,
		"title":               title,
		"update_tips":         updateTips,
	})
}
