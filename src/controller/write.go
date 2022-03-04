package controller

import (
	"github.com/gin-gonic/gin"
	"grayRelease/src/model"
	"grayRelease/src/service"
	"net/http"
)

func Config(c *gin.Context) {
	action := c.Param("action")
	if action == "" {
		// writes the given string into the response body.
		c.String(http.StatusBadRequest, "There is not a action named nil.")
		return
	} else if action == "add" {
		AddRule(c)
	} else if action == "delete" {
		DeleteRule(c)
	} else if action == "update" {
		UpdateRule(c)
	} else if action == "release" {
		ReleaseRule(c)
	} else if action == "offline" {
		OfflineRule(c)
	}
}

func getRule(c *gin.Context) *model.NewRule {
	var rule model.NewRule
	if err := c.ShouldBind(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	return &rule
}

func getId(c *gin.Context) uint {
	var id struct {
		Id uint `json:"id" form:"id"`
	}
	if err := c.ShouldBind(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	return id.Id
}

func AddRule(c *gin.Context) {
	rule := getRule(c)
	if service.CheckConfig(*rule) && service.AddRule(rule) {
		c.JSON(http.StatusOK, gin.H{"message": "the rule added successfully", "rule_aid": rule.Aid})
		return

	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "add rule fail"})
}

func DeleteRule(c *gin.Context) {
	aid := getId(c)
	if service.DeleteRule(aid) {
		c.JSON(http.StatusOK, gin.H{"message": "remove successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "remove fail"})
	}
	return
}

func UpdateRule(c *gin.Context) {
	rule := getRule(c)

	if service.CheckConfig(*rule) && service.UpdateRule(rule) {
		c.JSON(http.StatusOK, gin.H{"message": "please remember this aid", "rule_aid": rule.Aid})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "update fail"})
	}
}

func ReleaseRule(c *gin.Context) {
	aid := getId(c)
	if service.ReleaseRule(aid) {
		c.JSON(http.StatusOK, gin.H{"message": "release rule successfully", "rule_aid": aid})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "release rule fail", "rule_aid": aid})
}

func OfflineRule(c *gin.Context) {
	aid := getId(c)
	if service.OfflineRule(aid) {
		c.JSON(http.StatusOK, gin.H{"message": "offline rule successfully", "rule_aid": aid})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "offline rule fail", "rule_aid": aid})
}
