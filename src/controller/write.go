package controller

import (
	"github.com/gin-gonic/gin"
	"grayRelease/src/model"
	"grayRelease/src/service"
	"net/http"
	"strings"
)

func getRule(c *gin.Context) *model.Rule {
	var rule model.Rule
	if err := c.ShouldBind(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	return &rule
}

func getRuleId(c *gin.Context) uint {
	var id struct {
		Id uint `json:"id" form:"id"`
	}
	if err := c.ShouldBind(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	return id.Id
}

func AddRule(c *gin.Context) bool {
	rule := getRule(c)
	deviceIds := strings.Split(rule.DeviceIdList, ",")
	for i := 0; i < len(deviceIds); i++ {
		deviceIds[i] = strings.TrimSpace(deviceIds[i])
	}
	if service.CheckConfig(*rule) && service.AddRule(rule) {
		service.AddDeviceIdToWhiteList(int(rule.Id), deviceIds)
		c.JSON(http.StatusOK, gin.H{"message": "the rule added successfully", "rule_id": rule.Id})
		return true

	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "add rule fail"})
	return false
}

func DeleteRule(c *gin.Context) bool {
	id := getRuleId(c)
	if service.DeleteRule(id) {
		c.JSON(http.StatusOK, gin.H{"message": "remove successfully", "rule_id: ": id})
		return true
	}
	c.JSON(http.StatusOK, gin.H{"message": "remove fail", "rule_id: ": id})
	return false
}

func UpdateRule(c *gin.Context) bool {
	rule := getRule(c)

	if service.CheckConfig(*rule) && service.UpdateRule(rule) {
		c.JSON(http.StatusOK, gin.H{"message": "please remember this aid", "rule_id": rule.Id})
		return true
	}
	c.JSON(http.StatusOK, gin.H{"message": "update fail"})
	return false
}

func ReleaseRule(c *gin.Context) bool {
	ruleId := getRuleId(c)
	if service.ReleaseRule(ruleId) {
		c.JSON(http.StatusOK, gin.H{"message": "release rule successfully", "rule_id": ruleId})
		return true
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "release rule fail", "rule_aid": ruleId})
	return false
}

func OfflineRule(c *gin.Context) bool {
	ruleId := getRuleId(c)
	if service.OfflineRule(ruleId) {
		c.JSON(http.StatusOK, gin.H{"message": "offline rule successfully", "rule_id": ruleId})
		return true
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "offline rule fail", "rule_id": ruleId})
	return false
}
