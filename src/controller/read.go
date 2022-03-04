package controller

import (
	"github.com/gin-gonic/gin"
	"grayRelease/src/model"
	"grayRelease/src/service"
	"net/http"
)

func Check(c *gin.Context) {
	var client model.Client
	if err := c.ShouldBind(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	downloadUrl, updateVersionCode, md5, title, updateTips := service.Hit(&client)
	c.JSON(http.StatusOK, gin.H{
		"download_url:":       downloadUrl,
		"update_version_code": updateVersionCode,
		"md5":                 md5,
		"title":               title,
		"update_tips":         updateTips,
	})
}
