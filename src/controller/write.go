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
	}
}

func AddRule(c *gin.Context) {
	var rule model.Rule
	if err := c.ShouldBind(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if service.CheckConfig(rule) {
		if service.AddRule(&rule) {
			c.JSON(http.StatusOK, gin.H{"message": "the rule added successfully", "rule_aid": rule.Aid})
			return
		}
	}
	c.String(http.StatusBadRequest, "add rule fail")
}

func DeleteRule(c *gin.Context) {
	var id struct {
		Aid int `json:"aid"`
	}
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if service.DeleteRule(id.Aid) {
		c.JSON(http.StatusOK, gin.H{"message": "remove successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "remove fail"})
	}
	return
}

func UpdateRule(c *gin.Context) {
	var rule model.Rule
	if err := c.ShouldBind(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if service.CheckConfig(rule) {
		if service.AddRule(&rule) {
			c.JSON(http.StatusOK, gin.H{"message": "the rule update successfully", "rule_aid": rule.Aid})
			return
		}
	}

	if service.UpdateRule(&rule) {
		c.JSON(http.StatusOK, gin.H{"message": "please remember this aid", "rule_aid": rule.Aid})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "update fail"})
	}
}
