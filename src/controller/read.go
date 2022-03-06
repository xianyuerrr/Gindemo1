package controller

import (
	"github.com/gin-gonic/gin"
	"grayRelease/src/model/tables"
	"grayRelease/src/service"
	"net/http"
)

func Check(c *gin.Context) {
	var client tables.Client
	if err := c.ShouldBind(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := service.Hit(&client)
	downloadUrl := res[0]
	updateVersionCode := res[1]
	md5 := res[2]
	title := res[3]
	updateTips := res[4]
	c.JSON(http.StatusOK, gin.H{
		"download_url:":       downloadUrl,
		"update_version_code": updateVersionCode,
		"md5":                 md5,
		"title":               title,
		"update_tips":         updateTips,
	})
}
