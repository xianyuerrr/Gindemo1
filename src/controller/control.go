package controller

import (
	"demo1/src/model"
	"demo1/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Control() {
	//todo
}

func Config(c *gin.Context) {

	var config model.Rule
	if err := c.Bind(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if services.CheckConfig(config) {
		c.JSON(http.StatusOK, config)
	} else {
		c.String(http.StatusBadRequest, "check fail")
	}
}

func Check(c *gin.Context) {
	var form model.Client
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	downloadUrl, updateVersionCode, md5, title, updateTips := services.Hit(form)
	c.JSON(http.StatusOK, gin.H{
		"download_url:":       downloadUrl,
		"update_version_code": updateVersionCode,
		"md5":                 md5,
		"title":               title,
		"update_tips":         updateTips,
	})
}
